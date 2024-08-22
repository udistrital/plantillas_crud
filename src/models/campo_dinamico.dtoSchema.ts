import { ApiProperty} from '@nestjs/swagger';

export class CampoDinamicoDto{

    @ApiProperty()
    readonly nombre: string;

    @ApiProperty()
    readonly descripcion: string;
    
}

import { Schema, Prop, SchemaFactory } from "@nestjs/mongoose";
import { Document } from "mongoose";

@Schema({collection: 'campo_dinamico'})
export class CampoDinamico extends Document{

    @Prop({required: true})
    nombre: string

    @Prop({required: true})
    descripcion: string

}

export const CampoDinamicoSchema = SchemaFactory.createForClass(CampoDinamico);