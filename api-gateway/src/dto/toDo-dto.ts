import { ApiProperty } from "@nestjs/swagger"
import { IsOptional, IsString, Length } from "class-validator"

export class toDo {
    @ApiProperty({example: "Выпить чай", description: 'Название задачи'})
    @Length(1, 255, {message: 'От 1 до 255 символов'})
    @IsString({message: 'Должно быть строкой'})
    title:       string
    @ApiProperty({example: "...", description: 'Описание'})
    @Length(1, 255, {message: 'От 1 до 255 символов'})
    @IsString({message: 'Должно быть строкой'})
    @IsOptional()
    description: string
}