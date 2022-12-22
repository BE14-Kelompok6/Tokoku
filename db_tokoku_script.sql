use db_tokoku;
show tables;

DROP TABLE IF EXISTS role;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS barang;
DROP TABLE IF EXISTS transaksi;
DROP TABLE IF EXISTS pelanggan;
DROP TABLE IF EXISTS aktivitas_transaksi;


CREATE TABLE `role` (
`id` int AUTO_INCREMENT PRIMARY KEY,
`role` varchar(255) DEFAULT NULL
);

CREATE TABLE `users` (
`id` int AUTO_INCREMENT PRIMARY KEY,
`nama` varchar(255) DEFAULT NULL,
`password` varchar(255) DEFAULT NULL,
`role_id` int,
FOREIGN KEY (`role_id`) REFERENCES role (id)
);



CREATE TABLE `barang` (
id int AUTO_INCREMENT PRIMARY KEY,
user_id int NOT NULL,
nama_barang varchar(255) NOT NULL,
stok int NOT NULL,
tgl_input datetime NOT NULL,
FOREIGN KEY (`user_id`) REFERENCES users (id)
);

CREATE TABLE `pelanggan` (
id int AUTO_INCREMENT PRIMARY KEY,
user_id int ,
nama varchar(255) NOT NULL,
alamat varchar(255) NOT NULL,
FOREIGN KEY (`user_id`) REFERENCES users (id)
);


CREATE TABLE `transaksi` (
id int AUTO_INCREMENT PRIMARY KEY,
pelanggan_id int NOT NULL,
tgl_transaksi datetime NOT NULL,
pegawai int,
FOREIGN KEY (`pelanggan_id`) REFERENCES pelanggan (id),
FOREIGN KEY (`pegawai`) REFERENCES users (id)
);

CREATE TABLE `aktivitas_transaksi` (
barang_id int,
transaksi_id int,
qty int,
PRIMARY KEY (barang_id, transaksi_id),
CONSTRAINT fk_barang_transaksi_barang FOREIGN KEY (barang_id) REFERENCES barang (id),
CONSTRAINT fk_barang_transaksi_transaksi FOREIGN KEY (transaksi_id) REFERENCES transaksi(id)
);


INSERT INTO role VALUES(null, 'Admin');
INSERT INTO role VALUES(null, 'Pegawai');


INSERT INTO users  VALUES(null, 'admin', 'admin', 1);
SELECT * FROM users;
SELECT * FROM barang;

SELECT b.id, b.nama_barang, b.stok, b.tgl_input, u.nama  FROM barang b 
INNER JOIN users u 
ON u.id = b.user_id ;







