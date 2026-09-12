import { BusinessDayCalculatorEntityBase } from '../BusinessDayCalculatorEntityBase';
import type { BusinessDayCalculatorSDK } from '../BusinessDayCalculatorSDK';
import type { Control } from '../types';
import type { DnsResult, DnsResultLoadMatch } from '../BusinessDayCalculatorTypes';
declare class DnsResultEntity extends BusinessDayCalculatorEntityBase<DnsResult> {
    constructor(client: BusinessDayCalculatorSDK, entopts: any);
    make(this: DnsResultEntity): DnsResultEntity;
    load(this: any, reqmatch?: DnsResultLoadMatch, ctrl?: Control): Promise<DnsResultEntity>;
}
export { DnsResultEntity };
