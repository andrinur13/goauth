-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Mar 17, 2021 at 10:52 AM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 7.4.15

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_auth`
--

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(256) NOT NULL,
  `price` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `price`) VALUES
(1, 'Oreo Creamy', 6000),
(5, 'Oreo Choclate', 6500),
(6, 'Marimas Mangga', 300),
(7, 'Marimas Jeruk', 300),
(8, 'Permen Milkita', 2200),
(9, 'Minyak Tanah', 23000);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(256) NOT NULL,
  `name` varchar(256) NOT NULL,
  `email` varchar(256) NOT NULL,
  `password` varchar(512) NOT NULL,
  `role` varchar(10) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `name`, `email`, `password`, `role`, `created_at`, `updated_at`) VALUES
(1, 'andrinur13', 'Andri', 'andribis13@gmail.com', '$2a$04$kCpqpphZylF5uuHZ/zHqueJVQQpzMZ90Y9K8QiSMo6Z2WxtWaUJ3q', 'user', '2021-03-14 20:05:53', '2021-03-14 20:05:53'),
(2, 'ferhatm', 'Ferhat', 'ferhat@gmail.com', 'Ferhat', 'user', '2021-03-14 20:07:16', '2021-03-14 20:07:16'),
(3, 'user123', 'User 1 2 3', 'user123@gmail.com', '$2a$04$kCpqpphZylF5uuHZ/zHqueJVQQpzMZ90Y9K8QiSMo6Z2WxtWaUJ3q', 'user', '2021-03-16 10:37:42', '2021-03-16 10:37:42'),
(4, 'testing', 'Testing Uji Coba', 'testingtesting123@gmail.cm', '$2a$04$lh0KrQjqdFlY5zu4dGmn9ecSVaXiPZ3nKv8l/xAG4MCCYzyRhUnDG', 'user', '2021-03-16 11:19:03', '2021-03-16 11:19:03'),
(5, 'jono', 'Jono Joni', 'jonojoni@gmail.com', '$2a$04$o0OgAYyLd.5B0uYzHJdJJu/Bcx62OZtyyTW8CWrBzn1c1pnKuS41S', 'user', '2021-03-16 11:23:39', '2021-03-16 11:23:39'),
(6, 'bagas31', 'Bagas 31', 'bagas31@gmail.com', '$2a$04$gFDn5vf3kDo5PsOlrlvnYub/0zedpqidthVbelxhS4qr0lgcnRmxu', 'user', '2021-03-17 16:43:55', '2021-03-17 16:43:55');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
