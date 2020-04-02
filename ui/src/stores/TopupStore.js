import { EventEmitter } from "events";
import Swagger from "swagger-client";
import dispatcher from "../dispatcher";
import i18n, { packageNS } from '../i18n';
import { checkStatus, errorHandler } from "./helpers";
import sessionStore from "./SessionStore";




class TopupStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/topup.swagger.json", sessionStore.getClientOpts());
  }

  async getTopUpDestination(orgId) {
    console.log( 'getTopUpDestination',  orgId);
    try {
        const client = await this.swagger;
        let resp = await client.apis.TopUpService.GetTopUpDestination({
          orgId
        });
        
        resp = await checkStatus(resp);
        return resp.body;
      } catch (error) {
        errorHandler(error);
    }
  }

  getTopUpHistory(orgId, offset, limit, callbackFunc) {
    this.swagger.then(client => {
      client.apis.TopUpService.GetTopUpHistory({
        orgId,
        offset,
        limit
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.body);
      })
      .catch(errorHandler);
    });
  }

    getIncome(orgId, callbackFunc) {
    this.swagger.then(client => {
      client.apis.TopUpService.GetIncome({
        orgId
       
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.body);
      })
      .catch(errorHandler);
    });
  }

  
  notify(action) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "success",
        message: `${i18n.t(`${packageNS}:menu.store.transaction_has_been`)} ` + action,
      },
    });
  }
}

const topupStore = new TopupStore();
export default topupStore;
