import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { EventsModule } from './events/events.module';
import { PrismaModule } from './prisma/prisma.module';
import { SpotsModule } from './spots/spots.module';

@Module({
  imports: [ConfigModule.forRoot(), EventsModule, PrismaModule, SpotsModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
