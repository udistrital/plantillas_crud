import { Body, Controller, Delete, Get, HttpStatus, Param, Post, Put, Query, Res } from '@nestjs/common';
import { TipoPlantillaService } from '../services/tipo_plantilla.service';
import { TipoPlantillaDto } from '../models/tipo_plantilla.dtoSchema';
import { FilterDto } from 'src/filters/dto/filter.dto';
import { ApiTags } from '@nestjs/swagger';

@ApiTags('tipo_plantilla')
@Controller('tipo_plantilla')
export class TipoPlantillaController {
    constructor(private readonly tipoPlantillaService: TipoPlantillaService) {}

    @Post()
    async post(@Res() res, @Body() tipoPlantillaDto: TipoPlantillaDto) {
        this.tipoPlantillaService.post(tipoPlantillaDto).then(tipoPlantilla => {
            res.status(HttpStatus.CREATED).json({
                Success: true,
                Status: HttpStatus.CREATED,
                Message: 'Registration successful',
                Data: tipoPlantilla
            })
        }).catch(error => {
            res.status(HttpStatus.BAD_REQUEST).json({
                Success: false,
                Status: HttpStatus.BAD_REQUEST,
                Message: 'Error service Post: The request contains an incorrect data type or an invalid parameter',
                Data: error.message
            })
        });
    }

    @Get()
    async getAll(@Res() res, @Query() filterDto: FilterDto) {
        this.tipoPlantillaService.getAll(filterDto).then(tipoPlantilla => {
            res.status(HttpStatus.OK).json({
                Success: true,
                Status: HttpStatus.OK,
                Message: 'Request successful',
                Data: tipoPlantilla
            })
        }).catch(error => {
            res.status(HttpStatus.NOT_FOUND).json({
                Success: false,
                Status: HttpStatus.NOT_FOUND,
                Message: 'Error service GetAll: The request contains an incorrect parameter or no record exist',
                Data: error.message
            })
        })
    }

    @Get('/:id')
    async getById(@Res() res, @Param('id') id: string) {
        this.tipoPlantillaService.getById(id).then(tipoPlantilla => {
            res.status(HttpStatus.OK).json({
                Success: true,
                Status: HttpStatus.OK,
                Message: 'Request successful',
                Data: tipoPlantilla
            })
        }).catch(error => {
            res.status(HttpStatus.NOT_FOUND).json({
                Success: false,
                Status: HttpStatus.NOT_FOUND,
                Message: 'Error service GetOne: The request contains an incorrect parameter or no record exist',
                Data: error.message
            })
        })
    }

    @Put('/:id')
    async put(@Res() res, @Param('id') id: string, @Body() tipoPlantillaDto: TipoPlantillaDto) {
        this.tipoPlantillaService.put(id, tipoPlantillaDto).then(tipoPlantilla => {
            res.status(HttpStatus.OK).json({
                Success: true,
                Status: HttpStatus.OK,
                Message: 'Update successful',
                Data: tipoPlantilla
            })
        }).catch(error => {
            res.status(HttpStatus.BAD_REQUEST).json({
                Success: false,
                Status: HttpStatus.BAD_REQUEST,
                Message: 'Error service Put: The request contains an incorrect data type or an invalid parameter',
                Data: error.message
            })
        })
    }

    @Delete('/:id')
    async delete(@Res() res, @Param('id') id: string) {
        this.tipoPlantillaService.delete(id).then(tipoPlantilla => {
            res.status(HttpStatus.OK).json({
                Success: true,
                Status: HttpStatus.OK,
                Message: 'Delete successful',
                Data: tipoPlantilla
            })
        }).catch(error => {
            res.status(HttpStatus.NOT_FOUND).json({
                Success: false,
                Status: HttpStatus.NOT_FOUND,
                Message: 'Error service Delete: Request contains incorrect parameter',
                Data: error.message
            })
        })
    }
}
