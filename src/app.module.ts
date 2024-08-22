import { MiddlewareConsumer, Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { environment } from './config/configuration';
import { LoggerMiddleware } from './logger/logger';
import { CampoDinamico, CampoDinamicoSchema } from './models/campo_dinamico.dtoSchema';
import { CampoDinamicoController } from './controllers/campo_dinamico.controller';
import { CampoDinamicoService } from './services/campo_dinamico.service';
import { Plantilla, PlantillaSchema } from './models/plantilla.dtoSchema';
import { PlantillaController } from './controllers/plantilla.controller';
import { PlantillaService } from './services/plantilla.service';
import { Seccion, SeccionSchema } from './models/seccion.dtoSchema';
import { SeccionController } from './controllers/seccion.controller';
import { SeccionService } from './services/seccion.service';
import { TipoContenido, TipoContenidoSchema } from './models/tipo_contenido.dtoSchema';
import { TipoContenidoController } from './controllers/tipo_contenido.controller';
import { TipoContenidoService } from './services/tipo_contenido.service';
import { TipoPlantilla, TipoPlantillaSchema } from './models/tipo_plantilla.dtoSchema';
import { TipoPlantillaController } from './controllers/tipo_plantilla.controller';
import { TipoPlantillaService } from './services/tipo_plantilla.service';

@Module({
  imports: [
    MongooseModule.forRoot(`mongodb://localhost:27017`),
    // MongooseModule.forRoot(`mongodb://${environment.USER}:${environment.PASS}@`+`${environment.HOST}:${environment.PORT}/${environment.DB}?authSource=${environment.AUTH_DB}`),
    MongooseModule.forFeature([
      { name: CampoDinamico.name, schema: CampoDinamicoSchema },
      { name: Plantilla.name, schema: PlantillaSchema },
      { name: Seccion.name, schema: SeccionSchema },
      { name: TipoContenido.name, schema: TipoContenidoSchema },
      { name: TipoPlantilla.name, schema: TipoPlantillaSchema }
    ])
  ],
  controllers: [AppController, CampoDinamicoController, PlantillaController, SeccionController, TipoContenidoController, TipoPlantillaController],
  providers: [AppService, CampoDinamicoService, PlantillaService, SeccionService, TipoContenidoService, TipoPlantillaService],
})
export class AppModule {
  configure(consumer: MiddlewareConsumer): void {
    consumer.apply(LoggerMiddleware).forRoutes('*');
  }
}
