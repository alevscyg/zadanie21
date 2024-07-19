import { Body, Controller, Delete, Get, Inject, Param, Patch, Post, Req, UsePipes, ValidationPipe } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';
import { UserDto } from './dto/user-dto';
import { RoleDto } from './dto/role-dto';
import { ApiOperation, ApiResponse } from '@nestjs/swagger';
import { AuthDtoResponse } from './dto/auth-dto-response';
import { toDo } from './dto/toDo-dto';

@Controller("toDoList")
export class AppController {
  constructor(
    @Inject('RABBITMQ_AUTH_SERVICE') private rabbitAuthService: ClientProxy,
    @Inject('RABBITMQ_TODO_SERVICE') private rabbitToDoService: ClientProxy
  ) {}

  /// Авторизация ///////////////////////////////

  @UsePipes(ValidationPipe)
  @ApiOperation({summary: 'Регистрация'})
  @ApiResponse({status: 200, type: AuthDtoResponse})
  @Post("registration")
  async registration(
    @Body() userDto: UserDto) {
    return this.rabbitAuthService.send('registration', {...userDto});
  }

  @UsePipes(ValidationPipe)
  @ApiOperation({summary: 'Логин'})
  @ApiResponse({status: 200, type: AuthDtoResponse})
  @Post("login")
  async login(
    @Body() userDto: UserDto) {
    return this.rabbitAuthService.send('login', {...userDto});
  }

  /// Роли ///////////////////////////////

  @UsePipes(ValidationPipe)
  @ApiOperation({summary: 'Создание роли'})
  @ApiResponse({status: 200, type: RoleDto})
  @Post("role")
  async createRole(
    @Body() roleDto: RoleDto) {
    return this.rabbitAuthService.send('createRole', {...roleDto});
  }

  @UsePipes(ValidationPipe)
  @ApiOperation({summary: 'Обновить роль'})
  @ApiResponse({status: 200, type: RoleDto})
  @Patch("role")
  async updateRole(
    @Body() roleDto: RoleDto) {
    return this.rabbitAuthService.send('updateRole', {...roleDto});
  }

  @ApiOperation({summary: 'Получить все роли роли'})
  @ApiResponse({status: 200, type: RoleDto})
  @Get("role")
  async findAllRoles(){
    return this.rabbitAuthService.send('findAllRole', {});
  }

  @ApiOperation({summary: 'Найти роль'})
  @ApiResponse({status: 200, type: RoleDto})
  @Get("role/:value")
  async findOneRole(@Param('value') value: string,){
    return this.rabbitAuthService.send('findOneRole', { value });
  }

  @ApiOperation({summary: 'Удалить роль'})
  @ApiResponse({status: 200, type: RoleDto})
  @Delete("role/:value")
  async removeRole(@Param('value') value: string,){
    return this.rabbitAuthService.send('removeRole', { value });
  }

  /// ToDo ///////////////////////////////

  @UsePipes(ValidationPipe)
  @ApiOperation({summary: 'Создание ToDo'})
  @ApiResponse({status: 200, type: toDo})
  @Post("todo")
  async createToDo(
    @Body() toDoDto: toDo) {
    return this.rabbitToDoService.send('createToDo', {...toDoDto});
  }

  @ApiOperation({summary: 'Find ToDo by userId'})
  @ApiResponse({status: 200})
  @Get("todo/:valueSTR")
  async findToDoByUserId(@Param('valueSTR') valueSTR: string,){
    const value = Number(valueSTR)
    return this.rabbitToDoService.send('findToDoById', { value });
  }
  
}
