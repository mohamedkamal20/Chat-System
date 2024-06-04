#include <iostream>
#include <unordered_map>
#include <vector>
#include <mutex>
#include <chrono>
#include <ctime>

using namespace std;

class Account {
public:
    string accountId;
    string ownerName;
    double balance;
    
    Account() : accountId(""), ownerName(""), balance(0.0) {} // Default constructor
    Account(const string& id, const string& name, double initialBalance)
        : accountId(id), ownerName(name), balance(initialBalance) {}
};

class Transaction {
public:
    string transactionId;
    string fromAccountId;
    string toAccountId;
    double amount;
    time_t timestamp;

    Transaction(const string& id, const string& fromId, const string& toId, double amt)
        : transactionId(id), fromAccountId(fromId), toAccountId(toId), amount(amt) {
        timestamp = chrono::system_clock::to_time_t(chrono::system_clock::now());
    }
};

class PaymentProcessor {
private:
    unordered_map<string, Account> accounts;
    vector<Transaction> transactions;
    mutex mtx;
    int accountCounter = 0;
    int transactionCounter = 0;

    string generateAccountId() {
        return "ACC" + to_string(++accountCounter);
    }

    string generateTransactionId() {
        return "TXN" + to_string(++transactionCounter);
    }

public:
    void createAccount(const string& ownerName, double initialBalance) {
        if (initialBalance < 0) {
            cerr << "Error: Initial balance cannot be negative.\n";
            return;
        }
        string accountId = generateAccountId();
        lock_guard<mutex> lock(mtx);
        accounts[accountId] = Account(accountId, ownerName, initialBalance);
        cout << "Account created: " << accountId << ", Owner: " << ownerName << ", Balance: " << initialBalance << "\n";
    }

    bool processTransaction(const string& fromAccountId, const string& toAccountId, double amount) {
        lock_guard<mutex> lock(mtx);

        if (accounts.find(fromAccountId) == accounts.end() || accounts.find(toAccountId) == accounts.end()) {
            cout << "Transaction failed: Invalid account ID(s).\n";
            return false;
        }

        if (accounts[fromAccountId].balance < amount) {
            cout << "Transaction failed: Insufficient funds.\n";
            return false;
        }

        accounts[fromAccountId].balance -= amount;
        accounts[toAccountId].balance += amount;
        transactions.push_back(Transaction(generateTransactionId(), fromAccountId, toAccountId, amount));

        cout << "Transaction successful: " << fromAccountId << " -> " << toAccountId << ", Amount: " << amount << "\n";
        return true;
    }

    double getAccountBalance(const string& accountId) {
        lock_guard<mutex> lock(mtx);

        if (accounts.find(accountId) == accounts.end()) {
            cerr << "Error: Account ID not found.\n";
            return -1;
        }

        return accounts[accountId].balance;
    }
};

int main() {
    PaymentProcessor processor;

    // Create two new accounts
    processor.createAccount("Alice", 1000.0);
    processor.createAccount("Bob", 500.0);

    // Process a transaction
    processor.processTransaction("ACC1", "ACC2", 200.0);

    // Display current balances
    cout << "Balance of ACC1: " << processor.getAccountBalance("ACC1") << "\n";
    cout << "Balance of ACC2: " << processor.getAccountBalance("ACC2") << "\n";

    return 0;
}

