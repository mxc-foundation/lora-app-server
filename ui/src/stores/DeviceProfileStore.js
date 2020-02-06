import { EventEmitter } from "events";

import Swagger from "swagger-client";

import sessionStore from "./SessionStore";
import {checkStatus, errorHandler } from "./helpers";
import dispatcher from "../dispatcher";
import MockDeviceProfileStoreApi from '../api/mockDeviceProfileStoreApi';
import isDev from '../util/isDev';


class DeviceProfileStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/deviceProfile.swagger.json", sessionStore.getClientOpts())
  }

  create(deviceProfile, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceProfileService.Create({
        body: {
          deviceProfile: deviceProfile,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.notify("created");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  get(id, callbackFunc) {
    // Run the following in development environment and early exit from function
    /* if (isDev) {
      (async () => callbackFunc(await MockDeviceProfileStoreApi.get()))();
      return;
    } */

    this.swagger.then(client => {
      client.apis.DeviceProfileService.Get({
        id: id,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  update(deviceProfile, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceProfileService.Update({
        "deviceProfile.id": deviceProfile.id,
        body: {
          deviceProfile: deviceProfile,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.notify("updated");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  delete(id, callbackFunc) {
    this.swagger.then(client => {
      client.apis.DeviceProfileService.Delete({
        id: id,
      })
      .then(checkStatus)
      .then(resp => {
        this.notify("deleted");
        callbackFunc(resp.ojb);
      })
      .catch(errorHandler);
    });
  }

  list(organizationID, applicationID, limit, offset, callbackFunc) {
    // Run the following in development environment and early exit from function
    /* if (isDev) {
      (async () => callbackFunc(await MockDeviceProfileStoreApi.list()))();
      return;
    } */

    this.swagger.then(client => {
      client.apis.DeviceProfileService.List({
        organizationID: organizationID,
        applicationID: applicationID,
        limit: limit,
        offset: offset,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  notify(action) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "success",
        message: "device-profile has been " + action,
      },
    });
  }
}

const deviceProfileStore = new DeviceProfileStore();
export default deviceProfileStore;
