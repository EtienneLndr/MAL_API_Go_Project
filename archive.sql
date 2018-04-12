-- MySQL dump 10.16  Distrib 10.1.26-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: archive
-- ------------------------------------------------------
-- Server version	10.1.26-MariaDB-0+deb9u1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Archive`
--

DROP TABLE IF EXISTS `Archive`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Archive` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `objectInstanceIdentifier` bigint(20) unsigned DEFAULT NULL,
  `element` binary(20) DEFAULT NULL,
  `objectTypeArea` smallint(6) DEFAULT NULL,
  `objectTypeService` smallint(6) DEFAULT NULL,
  `objectTypeVersion` tinyint(4) DEFAULT NULL,
  `objectTypeNumber` smallint(6) DEFAULT NULL,
  `domain` text,
  `objectDetails` binary(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Archive`
--

LOCK TABLES `Archive` WRITE;
/*!40000 ALTER TABLE `Archive` DISABLE KEYS */;
INSERT INTO `Archive` VALUES (1,4857405852559299056,'\0\0\0\0\0\0\0\0\0\0\0\0\0',2,2,1,1,'/test/archiveService/','\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0test\0\0\0archiveService\0'),(2,5302008630720477549,'\0\0\0\0\0\r\0\0\0\0\0\0\0\0\0\0\0',2,2,1,13,'/test/archiveService/','\0\0\0\0\0\0\0\0\0\0\r\0\0\0\0\0\0test\0\0\0archiveService\0');
/*!40000 ALTER TABLE `Archive` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-04-11 13:49:53