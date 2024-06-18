/*
  Warnings:

  - The values [cancelled] on the enum `ReservationHistory_status` will be removed. If these variants are still used in the database, this will fail.
  - The values [unavailable] on the enum `Spot_status` will be removed. If these variants are still used in the database, this will fail.

*/
-- AlterTable
ALTER TABLE `ReservationHistory` MODIFY `status` ENUM('reserved', 'canceled') NOT NULL;

-- AlterTable
ALTER TABLE `Spot` MODIFY `status` ENUM('available', 'reserved') NOT NULL;
