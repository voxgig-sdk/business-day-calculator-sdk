import { BusinessDayCalculatorEntityBase } from '../BusinessDayCalculatorEntityBase';
import type { BusinessDayCalculatorSDK } from '../BusinessDayCalculatorSDK';
import type { Control } from '../types';
import type { Ipn, IpnLoadMatch } from '../BusinessDayCalculatorTypes';
declare class IpnEntity extends BusinessDayCalculatorEntityBase<Ipn> {
    constructor(client: BusinessDayCalculatorSDK, entopts: any);
    make(this: IpnEntity): IpnEntity;
    load(this: any, reqmatch?: IpnLoadMatch, ctrl?: Control): Promise<IpnEntity>;
}
export { IpnEntity };
