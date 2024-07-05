import { Controller } from '@nestjs/common';
import { MessagePattern, Payload } from '@nestjs/microservices';
import { AuthService } from './auth.service';
import { UserDto } from 'src/user/dto/user-dto';


@Controller()
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @MessagePattern('login')
  login(@Payload() authDto: UserDto) {
    return this.authService.login(authDto);
  }

  @MessagePattern('registration')
  registration(@Payload() authDto: UserDto) {
    return this.authService.registration(authDto);
  }
  
}
