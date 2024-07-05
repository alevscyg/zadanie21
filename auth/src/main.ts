import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ConfigService } from '@nestjs/config';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { DocumentBuilder, SwaggerModule } from "@nestjs/swagger";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const configService = app.get(ConfigService)
  const USER = configService.get('RABBITMQ_DEFAULT_USER');
  const PASSWORD =  configService.get('RABBITMQ_DEFAULT_PASS');
  const HOST = configService.get('RABBITMQ_HOST');
  const QUEUE = configService.get('RABBITMQ_AUTH_QUEUE');

  const config = new DocumentBuilder()
        .setTitle('Zadanie21')
        .setDescription('Документация REST API')
        .setVersion('1.0.0')
        .addTag('Zadanie21')
        .build()
    const document = SwaggerModule.createDocument(app, config);
    SwaggerModule.setup('/api/docs', app, document)

  app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.RMQ,
    options: {
      urls: [`amqp://${USER}:${PASSWORD}@${HOST}`],
      noAck:false,
      queue: QUEUE,
      queueOptions: {
        durable: true
      }
    }
  })
  await app.startAllMicroservices();
}
bootstrap();