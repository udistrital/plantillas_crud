import { ApiProperty} from '@nestjs/swagger';

export class PlantillaDto {

    @ApiProperty()
    readonly tipo_plantilla_id: string;

    @ApiProperty()
    secciones: any[];

    @ApiProperty()
    readonly sistema_id: string;

    @ApiProperty()
    grupo_id: string;

    @ApiProperty()
    version: string;

    @ApiProperty()
    nombre: string;

    @ApiProperty()
    uid: boolean;

    @ApiProperty()
    estado: boolean;

    @ApiProperty()
    activo: boolean;

    @ApiProperty()
    fecha_creacion: Date;

    @ApiProperty()
    fecha_modificacion: Date;
    
}

import { Schema, Prop, SchemaFactory } from "@nestjs/mongoose";
import { Document } from "mongoose";

@Schema({collection: 'plantilla'})
export class Plantilla extends Document {

    @Prop({required: true})
    tipo_plantilla_id: string;

    @Prop({required: true})
    secciones: any[];

    @Prop({required: true})
    sistema_id: string;

    @Prop({required: true})
    grupo_id: string;

    @Prop({required: true})
    version: string;

    @Prop({required: true})
    nombre: string;

    @Prop({required: true})
    uid: boolean;

    @Prop({required: true})
    estado: boolean;

    @Prop({required: true})
    activo: boolean;

    @Prop({required: true})
    fecha_creacion: Date;

    @Prop({required: true})
    fecha_modificacion: Date;

}

export const PlantillaSchema = SchemaFactory.createForClass(Plantilla);