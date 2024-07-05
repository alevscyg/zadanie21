import { Module } from '@nestjs/common';
import { UserModule } from './user/user.module';
import { ConfigModule } from '@nestjs/config';
import { AuthModule } from './auth/auth.module';
import { RoleModule } from './role/role.module';
import { DatabaseModule } from './database/database.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      envFilePath: `./.env`,
      isGlobal:true
    }),
    UserModule,
    AuthModule,
    RoleModule,
    DatabaseModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
