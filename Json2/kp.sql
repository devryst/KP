/*
SQLyog Enterprise v13.1.1 (64 bit)
MySQL - 10.4.13-MariaDB : Database - kp
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`kp` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `kp`;

/*Table structure for table `foto` */

DROP TABLE IF EXISTS `foto`;

CREATE TABLE `foto` (
  `Id` int(200) NOT NULL AUTO_INCREMENT,
  `Color` varchar(200) DEFAULT NULL,
  `Category` varchar(200) DEFAULT NULL,
  `Type` varchar(200) DEFAULT NULL,
  `Rgba` varchar(200) DEFAULT NULL,
  `Hex` varchar(200) DEFAULT NULL,
  `URL` varchar(200) DEFAULT NULL,
  `Width` varchar(200) DEFAULT NULL,
  `Height` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `Color` (`Color`),
  KEY `Category` (`Category`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

/*Data for the table `foto` */

/*Table structure for table `human` */

DROP TABLE IF EXISTS `human`;

CREATE TABLE `human` (
  `Id` int(200) NOT NULL AUTO_INCREMENT,
  `Age` int(200) DEFAULT NULL,
  `Detail` varchar(200) DEFAULT NULL,
  `Pointer` varchar(200) DEFAULT NULL,
  `Status` varchar(200) DEFAULT NULL,
  `Title` varchar(200) DEFAULT NULL,
  `Name` varchar(200) DEFAULT NULL,
  `Task1` varchar(200) DEFAULT NULL,
  `Task2` varchar(200) DEFAULT NULL,
  `Task3` varchar(200) DEFAULT NULL,
  `SecretIdentity` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `human` */

insert  into `human`(`Id`,`Age`,`Detail`,`Pointer`,`Status`,`Title`,`Name`,`Task1`,`Task2`,`Task3`,`SecretIdentity`) values 
(1,78,'hfh','yr','adad','addg','ada','j','g','h','WSGH'),
(2,29,'First name must contain at least three characters.','/data/attributes/firstName','422','Invalid Attribute','Molecule Man','Radiation resistance','Turning tiny','Radiation blast','Dan Jukes');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
