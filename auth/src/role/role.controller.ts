import { Controller } from '@nestjs/common';
import { MessagePattern, Payload } from '@nestjs/microservices';
import { RoleService } from './role.service';
import { RoleDto } from './dto/role-dto';


@Controller()
export class RoleController {
  constructor(private readonly roleService: RoleService) {}

  @MessagePattern('createRole')
  async create(@Payload() createRoleDto: RoleDto) {
    return await this.roleService.create(createRoleDto);
  }

  @MessagePattern('findAllRole')
  async findAll() {
    return await this.roleService.findAll();
  }

  @MessagePattern('findOneRole')
  async findOne(@Payload() value: object) {
    return this.roleService.findOne(value);
  }

  @MessagePattern('updateRole')
  async update(@Payload() updateRoleDto: RoleDto) {
    return this.roleService.update(updateRoleDto);
  }

  @MessagePattern('removeRole')
  async remove(@Payload() value: object) {
    return this.roleService.remove(value);
  }
}
