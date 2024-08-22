import { Injectable } from '@nestjs/common';
import { Model } from "mongoose";
import { InjectModel } from "@nestjs/mongoose";
import { FiltersService } from 'src/filters/filters.service';
import { FilterDto } from 'src/filters/dto/filter.dto';
import { SeccionDto as MainDto, Seccion as MainModel } from '../models/seccion.dtoSchema';
import { TipoContenido } from 'src/models/tipo_contenido.dtoSchema';
import { CampoDinamico } from 'src/models/campo_dinamico.dtoSchema';

@Injectable()
export class SeccionService {
    constructor(
        @InjectModel(MainModel.name)
        private readonly mainModel: Model<MainModel>,

        // ? inyectar modelos relacionados para verificar existencia y popular
        @InjectModel(TipoContenido.name)
        private readonly tipoContenidoModel: Model<TipoContenido>,
        @InjectModel(CampoDinamico.name)
        private readonly campoDinamicoModel: Model<CampoDinamico>
    ) {}

    /**
     * Revisa si los _ids de las colecciones relacionadas existen.
     * Detiene el post o put si no hay concordancia.
     * ? Agregar demás comprobaciones si se añaden más relaciones.
     * @param mainDto - modelo dto que se desea verificar.
     */
    private async checkRelated(mainDto: MainDto) { 
        if (mainDto.padre_id && mainDto.padre_id != null) {
            const padreId = await this.mainModel.exists({ _id: mainDto.padre_id });
            if (!padreId) {
                throw new Error(`padre_id: ${mainDto.padre_id} doesn't exist`);
            }
        }
        if (mainDto.tipo_contenido_id) {
            const tipoContenidoId = await this.tipoContenidoModel.exists({ _id: mainDto.tipo_contenido_id });
            if (!tipoContenidoId) {
                throw new Error(`tipo_contenido_id: ${mainDto.tipo_contenido_id} doesn't exist`);
            }
        }
        if (mainDto.campos_dinamicos) {
            for (const campo_dinamico of mainDto.campos_dinamicos) {
                const campoDinamicoId = await this.campoDinamicoModel.exists({ _id: campo_dinamico });
                if (!campoDinamicoId) {
                    throw new Error(`campo_dinamico: ${campo_dinamico} doesn't exist`);
                }
            }
        }
    }

    /**
     * Retorna la lista de colecciones a popular segun relación con la coleción actual.
     * ? Agregar aquí si se relacionan más colecciones.
     */
    private populatefields(): any[] {
        return [
            { path: TipoContenido.name + 'Id' },
            { path: CampoDinamico.name + 'Id' }
        ]
    }

    // ? funciones REST generalizadas
    async post(mainDto: MainDto): Promise<MainDto> {
        const dateNow = new Date();
        const newdoc = new this.mainModel(mainDto);
        newdoc.fecha_creacion = dateNow;
        newdoc.fecha_modificacion = dateNow;
        await this.checkRelated(newdoc);
        return await newdoc.save();
    }

    async getAll(filterDto: FilterDto): Promise<MainDto[]> {
        const filterService = new FiltersService(filterDto);
        let populatefields = [];
        if (filterService.isPopulated()) {
            populatefields = this.populatefields();
        }
        return await this.mainModel.find(
            filterService.getQuery(),
            filterService.getFields(),
            filterService.getLimitAndOffset()
        ).sort(
            filterService.getSortBy()
        ).populate(populatefields);
    }

    async getById(_id: string): Promise<MainDto> {
        const doc = await this.mainModel.findById(_id);
        if (!doc) {
            throw new Error(`${_id} doesn't exist`);
        }
        return doc;
    }

    async put(_id: string, mainDto: MainDto): Promise<MainDto> {
        mainDto.fecha_modificacion = new Date();
        if (mainDto.fecha_creacion) {
            delete mainDto.fecha_creacion;
        }
        await this.checkRelated(mainDto);
        return await this.mainModel.findByIdAndUpdate(_id, mainDto, { new: true });
    }

    async delete(_id: string): Promise<MainDto> {
        //const deleted = await this.mainModel.findByIdAndDelete(_id);
        const deleted = await this.mainModel.findByIdAndUpdate(_id, { activo: false }, { new: true });
        if (!deleted) {
            throw new Error(`${_id} doesn't exist`);
        }
        return deleted;
    }
}
