/*
SQLyog Ultimate v12.4.3 (64 bit)
MySQL - 5.6.21 : Database - northwind
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`northwind` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `northwind`;

/*Table structure for table `employees` */

DROP TABLE IF EXISTS `employees`;

CREATE TABLE `employees` (
  `EmployeeID` int(11) NOT NULL AUTO_INCREMENT,
  `LastName` varchar(20) DEFAULT NULL,
  `FirstName` varchar(10) DEFAULT NULL,
  `Title` varchar(30) DEFAULT NULL,
  `TitleOfCourtesy` varchar(25) DEFAULT NULL,
  `BirthDate` date DEFAULT NULL,
  `HireDate` date DEFAULT NULL,
  `Address` varchar(60) DEFAULT NULL,
  `City` varchar(15) NOT NULL,
  `Region` varchar(15) DEFAULT NULL,
  `PostalCode` varchar(10) DEFAULT NULL,
  `Country` varchar(15) DEFAULT NULL,
  `HomePhone` varchar(24) DEFAULT NULL,
  `Extension` varchar(4) DEFAULT NULL,
  `Photo` varchar(255) DEFAULT NULL,
  `Notes` text,
  `ReportsTo` int(11) DEFAULT NULL,
  `ProvinceName` varchar(75) DEFAULT NULL,
  PRIMARY KEY (`EmployeeID`),
  KEY `LastName` (`LastName`),
  KEY `PostalCode` (`PostalCode`)
) ENGINE=MyISAM AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;

/*Data for the table `employees` */

insert  into `employees`(`EmployeeID`,`LastName`,`FirstName`,`Title`,`TitleOfCourtesy`,`BirthDate`,`HireDate`,`Address`,`City`,`Region`,`PostalCode`,`Country`,`HomePhone`,`Extension`,`Photo`,`Notes`,`ReportsTo`,`ProvinceName`) values 
(1,'Davolio','Nancy','Sales Representative','Ms.','1971-01-31','1992-05-01','507 - 20th Ave. E.\r\nApt. 2A','Seattle','WA','98122','USA','(206) 555-9857','5467','EmpID1.bmp','Education includes a BA in psychology from Colorado State University.  She also completed \"The Art of the Cold Call.\"  Nancy is a member of Toastmasters International.',2,NULL),
(3,'Leverling','Janet','Sales Representative','Ms.','1963-08-30','1992-04-01','722 Moss Bay Blvd.','Kirkland','WA','98033','USA','(206) 555-3412','3355','EmpID3.bmp','Janet has a BS degree in chemistry from Boston College).  She has also completed a certificate program in food retailing management.  Janet was hired as a sales associate and was promoted to sales representative.',2,NULL),
(4,'Peacock','Margaretss','Sales Representative','Mrs.','1970-01-29','1993-05-03','4110 Old Redmond Rd.','Redmond','WA','98052','USA','(206) 555-8122','5176','EmpID4.bmp','Margaret holds a BA in English literature from Concordia College and an MA from the American Institute of Culinary Arts. She was temporarily assigned to the London office before returning to her permanent post in Seattle.',2,NULL),
(5,'Buchanan','Steven','Sales Manager','Mr.','1977-01-21','1993-10-17','14 Garrett Hill','London',NULL,'SW1 8JR','UK','(71) 555-4848','3453','EmpID5.bmp','Steven Buchanan graduated from St. Andrews University, Scotland, with a BSC degree.  Upon joining the company as a sales representative, he spent 6 months in an orientation program at the Seattle office and then returned to his permanent post in London, where he was promoted to sales manager.  Mr. Buchanan has completed the courses \"Successful Telemarketing\" and \"International Sales Management.\"  He is fluent in French.',2,NULL),
(7,'King','Robert','Sales Representative','Mr.','1970-01-07','1994-01-02','Edgeham Hollow\r\nWinchester Way','London',NULL,'RG1 9SP','UK','(71) 555-5598','465','EmpID7.bmp','Robert King served in the Peace Corps and traveled extensively before completing his degree in English at the University of Michigan and then joining the company.  After completing a course entitled \"Selling in Europe,\" he was transferred to the London office.',5,NULL),
(8,'Callahan','Laura','Inside Sales Coordinator','Ms.','1970-01-01','1994-03-05','4726 - 11th Ave. N.E.','Seattle','WA','98105','USA','(206) 555-1189','2344','EmpID8.bmp','Laura received a BA in psychology from the University of Washington.  She has also completed a course in business French.  She reads and writes French.',2,NULL),
(9,'Dodsworth','Anne','Sales Representative','Ms.','1970-01-01','1994-11-15','7 Houndstooth Rd.','London',NULL,'WG2 7LT','UK','(71) 555-4444','452','EmpID9.bmp','Anne has a BA degree in English from St. Lawrence College.  She is fluent in French and German.',5,NULL),
(10,'tettt','tet',NULL,NULL,'1970-01-01',NULL,NULL,'',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL),
(11,'testdf','testdf',NULL,NULL,NULL,NULL,NULL,'',NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL,NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
