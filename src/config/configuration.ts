import { config } from "dotenv";
config()

export const environment =  {
    USER: process.env.PLANTILLAS_CRUD_USER,
    PASS: process.env.PLANTILLAS_CRUD_PASS,
    HOST: process.env.PLANTILLAS_CRUD_HOST,
    PORT: process.env.PLANTILLAS_CRUD_PORT,
    DB: process.env.PLANTILLAS_CRUD_DB,
    HTTP_PORT: process.env.PLANTILLAS_CRUD_HTTP_PORT,
    AUTH_DB: process.env.PLANTILLAS_CRUD_AUTH_DB 
};