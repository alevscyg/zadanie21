import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ConfigService } from '@nestjs/config';
import { DocumentBuilder, SwaggerModule } from "@nestjs/swagger";

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {cors:true});
  const configService = app.get(ConfigService)
  const PORT = configService.get('PORT');
  const config = new DocumentBuilder()
        .setTitle('Zadanie21')
        .setDescription('Документация REST API')
        .setVersion('1.0.0')
        .addTag('Zadanie21')
        .build()
    const document = SwaggerModule.createDocument(app, config);
    SwaggerModule.setup('/api/docs', app, document)

  await app.listen(PORT, () => console.log(`API Gateway started on port = ${PORT}`)) 
}
bootstrap();
