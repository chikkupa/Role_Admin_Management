-- phpMyAdmin SQL Dump
-- version 4.6.6deb5
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 10, 2018 at 10:58 PM
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

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `username` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `last_name` varchar(50) DEFAULT '',
  `email` varchar(100) DEFAULT NULL,
  `dob` varchar(30) DEFAULT '',
  `gender` varchar(10) DEFAULT '',
  `race` varchar(20) DEFAULT '',
  `street` varchar(255) DEFAULT '',
  `city` varchar(50) DEFAULT '',
  `phone` varchar(20) DEFAULT '',
  `password` varchar(255) DEFAULT NULL,
  `role` int(11) DEFAULT '2',
  `status` int(11) DEFAULT '0',
  `verify_string` varchar(255) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `username`, `name`, `last_name`, `email`, `dob`, `gender`, `race`, `street`, `city`, `phone`, `password`, `role`, `status`, `verify_string`, `create_time`) VALUES
(3, 'admin', 'Administrator', 'Last Name', 'admin@user.com', '', 'Male', '', '', '', '', '40bd001563085fc35165329ea1ff5c5ecbdbbeef', 1, 1, NULL, '2018-07-04 16:31:07'),
(4, 'user_1', 'User Type 1', 'Last Name', 'user1@user.com', '', 'Male', '', '', '', '', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 2, 1, NULL, '2018-07-05 18:33:42'),
(5, 'user_2', 'User Type 2', 'Last Name', 'user2@user.com', '', 'Male', '', '', '', '', '8cb2237d0679ca88db6464eac60da96345513964', 3, 1, NULL, '2018-07-05 18:47:12'),
(6, 'user_3', 'User Type 3', 'Last Name', 'user3@user.com', '', 'Male', '', '', '', '', '40bd001563085fc35165329ea1ff5c5ecbdbbeef', 4, 1, NULL, '2018-07-06 09:54:09');

-- --------------------------------------------------------

--
-- Table structure for table `user_files`
--

DROP TABLE IF EXISTS `user_files`;
CREATE TABLE `user_files` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `file_category` varchar(50) DEFAULT NULL,
  `filename` text
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
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

-- --------------------------------------------------------

--
-- Table structure for table `user_status`
--

DROP TABLE IF EXISTS `user_status`;
CREATE TABLE `user_status` (
  `id` int(11) NOT NULL,
  `status` varchar(50) DEFAULT NULL,
  `value` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `user_status`
--

INSERT INTO `user_status` (`id`, `status`, `value`) VALUES
(1, 'Pending Email Verification', 0),
(2, 'Reserved Status', 1),
(3, 'Pendent', 2),
(4, 'OK', 3),
(5, 'Uploaded', 4),
(6, 'Verified', 5),
(7, 'Approved', 6),
(8, 'Meet', 7),
(9, 'Accepted', 8),
(10, 'Denied', 9);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_files`
--
ALTER TABLE `user_files`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_roles`
--
ALTER TABLE `user_roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_status`
--
ALTER TABLE `user_status`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;
--
-- AUTO_INCREMENT for table `user_files`
--
ALTER TABLE `user_files`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;
--
-- AUTO_INCREMENT for table `user_roles`
--
ALTER TABLE `user_roles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `user_status`
--
ALTER TABLE `user_status`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
