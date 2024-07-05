import { Injectable } from '@nestjs/common';
import { DatabaseService } from 'src/database/database.service';
import { RoleDto } from './dto/role-dto';


@Injectable()
export class RoleService {
      
  constructor(private databaseService: DatabaseService){}

  async create(createRoleDto: RoleDto) {
    return await this.databaseService.role.create({ data: createRoleDto });
  }
    
  async findAll() {
    return await this.databaseService.role.findMany({});
  }
    
  async findOne(value: object) {
    // @ts-expect-error
    return await this.databaseService.role.findUnique({ where: {value: value.value}})
  }
  
  async update(updateRoleDto: RoleDto) {
    return await this.databaseService.role.update({
      where: {
        value: updateRoleDto.value,
      },
      data: updateRoleDto,
    });
  }
    
  async remove(value: object) {
    // @ts-expect-error
    return await this.databaseService.role.delete({ where: {value: value.value}});
  }
  
}
