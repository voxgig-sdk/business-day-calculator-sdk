import { BusinessDayCalculatorEntityBase } from '../BusinessDayCalculatorEntityBase';
import type { BusinessDayCalculatorSDK } from '../BusinessDayCalculatorSDK';
import type { Control } from '../types';
import type { Generate, GenerateLoadMatch } from '../BusinessDayCalculatorTypes';
declare class GenerateEntity extends BusinessDayCalculatorEntityBase<Generate> {
    constructor(client: BusinessDayCalculatorSDK, entopts: any);
    make(this: GenerateEntity): GenerateEntity;
    load(this: any, reqmatch?: GenerateLoadMatch, ctrl?: Control): Promise<GenerateEntity>;
}
export { GenerateEntity };
