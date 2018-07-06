-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 06, 2018 at 08:29 PM
-- Server version: 5.7.22-0ubuntu18.04.1
-- PHP Version: 7.0.30-1+ubuntu18.04.1+deb.sury.org+1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_login_app`
--

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `username` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT '',
  `email` varchar(100) DEFAULT NULL,
  `dob` varchar(30) DEFAULT '',
  `gender` varchar(10) DEFAULT '',
  `race` varchar(20) DEFAULT NULL,
  `street` varchar(255) DEFAULT NULL,
  `city` varchar(50) DEFAULT '',
  `phone` varchar(20) DEFAULT '',
  `password` varchar(255) DEFAULT NULL,
  `role` int(11) DEFAULT '2',
  `status` int(11) DEFAULT '2',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `username`, `name`, `last_name`, `email`, `dob`, `gender`, `race`, `street`, `city`, `phone`, `password`, `role`, `status`, `create_time`) VALUES
(3, 'admin', 'Administrator', 'Last Name', 'admin@user.com', '', 'Male', NULL, NULL, '', '', '40bd001563085fc35165329ea1ff5c5ecbdbbeef', 1, 1, '2018-07-04 16:31:07'),
(4, 'user_1', 'User Type 1', 'Last Name', 'user1@user.com', '', 'Male', NULL, NULL, '', '', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 2, 1, '2018-07-05 18:33:42'),
(5, 'user_2', 'User Type 2', 'Last Name', 'user2@user.com', '', 'Male', NULL, NULL, '', '', '8cb2237d0679ca88db6464eac60da96345513964', 3, 1, '2018-07-05 18:47:12'),
(6, 'user_3', 'User Type 3', 'Last Name', 'user3@user.com', '', 'Male', NULL, NULL, '', '', '40bd001563085fc35165329ea1ff5c5ecbdbbeef', 4, 1, '2018-07-06 09:54:09');

-- --------------------------------------------------------

--
-- Table structure for table `user_roles`
--

CREATE TABLE `user_roles` (
  `id` int(11) NOT NULL,
  `role` varchar(50) DEFAULT NULL,
  `value` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `user_roles`
--

INSERT INTO `user_roles` (`id`, `role`, `value`) VALUES
(1, 'admin', 1),
(2, 'user_type_1', 2),
(3, 'user_type_2', 3),
(4, 'user_type_3', 4);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_roles`
--
ALTER TABLE `user_roles`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
--
-- AUTO_INCREMENT for table `user_roles`
--
ALTER TABLE `user_roles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
