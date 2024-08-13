--Menyimpan informasi pengguna.
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Menyimpan informasi dompet yang dimiliki pengguna.
CREATE TABLE wallets (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    balance DECIMAL(15, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Menyimpan kategori transaksi yang dibuat oleh pengguna.
CREATE TABLE transaction_categories (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Menyimpan informasi transaksi pemasukan/pengeluaran yang dilakukan pengguna dalam dompet tertentu.
CREATE TABLE records (
    id SERIAL PRIMARY KEY,
    wallet_id INT REFERENCES wallets(id) ON DELETE CASCADE,
    category_id INT REFERENCES transaction_categories(id),
    amount DECIMAL(15, 2) NOT NULL,
    type VARCHAR(10) CHECK (type IN ('income', 'expense')) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Menyimpan informasi transfer antara dua dompet.
CREATE TABLE transfers (
    id SERIAL PRIMARY KEY,
    from_wallet_id INT REFERENCES wallets(id) ON DELETE CASCADE,
    to_wallet_id INT REFERENCES wallets(id) ON DELETE CASCADE,
    amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


--Menampilkan Record Berdasarkan start_time dan end_time
SELECT * FROM records 
WHERE wallet_id = :wallet_id 
AND created_at BETWEEN :start_time AND :end_time;



-- Menampilkan Rekapitulasi Total Pemasukan dan Pengeluaran
SELECT 
    SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END) AS total_income,
    SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END) AS total_expense
FROM records
WHERE created_at BETWEEN :start_time AND :end_time;


-- Menampilkan Data Pengeluaran per Kategori
SELECT 
    tc.name AS category,
    SUM(r.amount) AS total_expense
FROM records r
JOIN transaction_categories tc ON r.category_id = tc.id
WHERE r.type = 'expense'
AND r.created_at BETWEEN :start_time AND :end_time
GROUP BY tc.name;


-- Menampilkan 10 Record Terakhir dari Semua Wallet
SELECT * FROM records 
WHERE wallet_id IN (SELECT id FROM wallets WHERE user_id = :user_id)
ORDER BY created_at DESC 
LIMIT 10;
