import { ApiProperty} from '@nestjs/swagger';

export class TipoContenidoDto {

    @ApiProperty()
    readonly nombre: string;

    @ApiProperty()
    readonly descripcion: string;

    @ApiProperty()
    readonly codigo_abreviacion: string;
    
}

import { Schema, Prop, SchemaFactory } from "@nestjs/mongoose";
import { Document } from "mongoose";

@Schema({collection: 'tipo_contenido'})
export class TipoContenido extends Document {

    @Prop({required: true})
    nombre: string;

    @Prop({required: true})
    descripcion: string;

    @Prop({required: true})
    codigo_abreviacion: string;

}

export const TipoContenidoSchema = SchemaFactory.createForClass(TipoContenido);