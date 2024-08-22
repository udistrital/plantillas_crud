import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export class SeccionDto {

    @ApiPropertyOptional({ type: String, nullable: true })
    readonly padre_id: string | null;

    @ApiProperty()
    readonly tipo_contenido_id: string;

    @ApiProperty()
    campos_dinamicos: any[];

    @ApiProperty()
    contenido: string;

    @ApiProperty()
    estilos: any[];

    @ApiProperty()
    orden: string;

    @ApiProperty()
    activo: boolean;

    @ApiProperty()
    fecha_creacion: Date;

    @ApiProperty()
    fecha_modificacion: Date;

}

import { Schema, Prop, SchemaFactory } from "@nestjs/mongoose";
import { Document } from "mongoose";

@Schema({collection: 'seccion'})
export class Seccion extends Document {

    @Prop({required: false, type: String})
    padre_id: string | null;

    @Prop({required: true})
    tipo_contenido_id: string;

    @Prop({required: true})
    campos_dinamicos: any[];

    @Prop({required: true})
    contenido: string;

    @Prop({required: true})
    estilos: any[];

    @Prop({required: true})
    orden: string;

    @Prop({required: true})
    activo: boolean

    @Prop({required: true})
    fecha_creacion: Date

    @Prop({required: true})
    fecha_modificacion: Date
    
}

export const SeccionSchema = SchemaFactory.createForClass(Seccion);
