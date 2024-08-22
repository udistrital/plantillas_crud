import { Body, Controller, Delete, Get, HttpStatus, Param, Post, Put, Query, Res } from '@nestjs/common';
import { CampoDinamicoService } from '../services/campo_dinamico.service';
import { CampoDinamicoDto } from '../models/campo_dinamico.dtoSchema';
import { FilterDto } from 'src/filters/dto/filter.dto';
import { ApiTags } from '@nestjs/swagger';

@ApiTags('campo_dinamico')
@Controller('campo_dinamico')
export class CampoDinamicoController {
    constructor(private readonly campoDinamicoService: CampoDinamicoService) {}

    @Post()
    async post(@Res() res, @Body() campoDinamicoDto: CampoDinamicoDto) {
        this.campoDinamicoService.post(campoDinamicoDto).then(campoDinamico => {
            res.status(HttpStatus.CREATED).json({
                Success: true,
                Status: HttpStatus.CREATED,
                Message: 'Registration successful',
                Data: campoDinamico
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
        this.campoDinamicoService.getAll(filterDto).then(campoDinamico => {
            res.status(HttpStatus.OK).json({
                Success: true,
                Status: HttpStatus.OK,
                Message: 'Request successful',
                Data: campoDinamico
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
        this.campoDinamicoService.getById(id).then(campoDinamico => {
            res.status(HttpStatus.OK).json({
                Success: true,
                Status: HttpStatus.OK,
                Message: 'Request successful',
                Data: campoDinamico
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
    async put(@Res() res, @Param('id') id: string, @Body() campoDinamicoDto: CampoDinamicoDto) {
        this.campoDinamicoService.put(id, campoDinamicoDto).then(campoDinamico => {
            res.status(HttpStatus.OK).json({
                Success: true,
                Status: HttpStatus.OK,
                Message: 'Update successful',
                Data: campoDinamico
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
        this.campoDinamicoService.delete(id).then(campoDinamico => {
            res.status(HttpStatus.OK).json({
                Success: true,
                Status: HttpStatus.OK,
                Message: 'Delete successful',
                Data: campoDinamico
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
