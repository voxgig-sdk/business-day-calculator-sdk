import { BusinessDayCalculatorEntityBase } from '../BusinessDayCalculatorEntityBase';
import type { BusinessDayCalculatorSDK } from '../BusinessDayCalculatorSDK';
import type { Control } from '../types';
import type { EmailValidate, EmailValidateLoadMatch } from '../BusinessDayCalculatorTypes';
declare class EmailValidateEntity extends BusinessDayCalculatorEntityBase<EmailValidate> {
    constructor(client: BusinessDayCalculatorSDK, entopts: any);
    make(this: EmailValidateEntity): EmailValidateEntity;
    load(this: any, reqmatch?: EmailValidateLoadMatch, ctrl?: Control): Promise<EmailValidateEntity>;
}
export { EmailValidateEntity };
