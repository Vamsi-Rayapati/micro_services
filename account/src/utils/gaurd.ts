import { NextFunction, Request, Response } from "express";
import { UniqueConstraintError, ValidationError as SqlValidationError } from "sequelize";
import { BaseError, ValidationDetail, ValidationError } from "../api/v1/models/error_models";

function handleError(error: any,next: NextFunction) {
    try {
        if(error instanceof UniqueConstraintError) {
            const sqlError = error.errors[0];
            const details:ValidationDetail[]= [{
                property: sqlError.path ?? 'data',
                reason: sqlError.message}]
            throw new BaseError(sqlError.message,409,details);
        } else if(error instanceof SqlValidationError) {
            const details:ValidationDetail[]= error.errors.map((sqlError)=> ({
                property: sqlError.path ?? 'data',
                reason: sqlError.message
            }))

            throw new ValidationError('Input Invalidation Error', details);
        } else {
            throw new BaseError("Internal Server Error", 500);
        }
    } catch(err) {
        next(err);
    }
}

function gaurd(callback: Function) {
    return async (req:Request,res:Response,next: NextFunction) => {
        try {
            await callback(req,res,next)
        } catch(err) {
            handleError(err, next)
        }
    }
  
}

export default gaurd;