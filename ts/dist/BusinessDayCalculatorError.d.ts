import { Context } from './Context';
declare class BusinessDayCalculatorError extends Error {
    isBusinessDayCalculatorError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { BusinessDayCalculatorError };
