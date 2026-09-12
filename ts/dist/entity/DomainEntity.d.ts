import { BusinessDayCalculatorEntityBase } from '../BusinessDayCalculatorEntityBase';
import type { BusinessDayCalculatorSDK } from '../BusinessDayCalculatorSDK';
import type { Control } from '../types';
import type { Domain, DomainListMatch } from '../BusinessDayCalculatorTypes';
declare class DomainEntity extends BusinessDayCalculatorEntityBase<Domain> {
    constructor(client: BusinessDayCalculatorSDK, entopts: any);
    make(this: DomainEntity): DomainEntity;
    list(this: any, reqmatch?: DomainListMatch, ctrl?: Control): Promise<DomainEntity[]>;
}
export { DomainEntity };
