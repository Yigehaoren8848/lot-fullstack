import fetch from '../config/fetch'
import {getStore} from '../config/mUtils'

/**
 * 设备控制
 */

export const controlEqui = (order) => fetch("/equipments/control", {
	order: order
});

