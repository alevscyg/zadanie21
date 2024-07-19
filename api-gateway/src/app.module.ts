import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { ClientProxyFactory, ClientsModule, Transport } from '@nestjs/microservices';

@Module({
  imports: [
    
    ConfigModule.forRoot({
      envFilePath: `./.development.env`,
    }),
  ],
  controllers: [AppController],
  providers: [AppService,
    {
      provide: 'RABBITMQ_AUTH_SERVICE',
        useFactory:(configService: ConfigService) => {
          const USER = configService.get('RABBITMQ_DEFAULT_USER');
          const PASSWORD =  configService.get('RABBITMQ_DEFAULT_PASS');
          const HOST = configService.get('RABBITMQ_HOST');
          const QUEUE = configService.get('RABBITMQ_AUTH_QUEUE');
    
          return ClientProxyFactory.create({
            transport: Transport.RMQ,
            options: {
              urls:[`amqp://${USER}:${PASSWORD}@${HOST}`],
              queue: QUEUE,
              queueOptions: {
                durable: false,
                noAck: true
              }
            }
          })
        },
        inject:[ConfigService]
    },
    {
      provide: 'RABBITMQ_TODO_SERVICE',
        useFactory:(configService: ConfigService) => {
          const USER = configService.get('RABBITMQ_DEFAULT_USER');
          const PASSWORD =  configService.get('RABBITMQ_DEFAULT_PASS');
          const HOST = configService.get('RABBITMQ_HOST');
          const QUEUE = configService.get('RABBITMQ_TODO_QUEUE');
    
          return ClientProxyFactory.create({
            transport: Transport.RMQ,
            options: {
              urls:[`amqp://${USER}:${PASSWORD}@${HOST}`],
              queue: QUEUE,
              queueOptions: {
                durable: false,
                noAck: true
              }
            }
          })
        },
        inject:[ConfigService]
    }
  ],
})
export class AppModule {}
