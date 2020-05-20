import { EventEmitter } from "events";
import Swagger from "swagger-client";
import dispatcher from "../dispatcher";
import { checkStatus, errorHandler } from "./helpers";
import sessionStore from "./SessionStore";




class ApplicationStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/application.swagger.json", sessionStore.getClientOpts());
  }

  create(application, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.Create({
        body: {
          application: application,
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
    this.swagger.then(client => {
      client.apis.ApplicationService.Get({
        id: id,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  update(application, callbackFunc, errorCallbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.Update({
        "application.id": application.id,
        body: {
          application: application,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.notify("updated");
        callbackFunc(resp.obj);
      })
      .catch(error => {
        errorHandler(error);
        if (errorCallbackFunc) errorCallbackFunc(error);
      });
    });
  }

  delete(id, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.Delete({
        id: id,
      })
      .then(checkStatus)
      .then(resp => {
        this.notify("deleted");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  list(search, organizationID, limit, offset, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.List({
        limit: limit,
        offset: offset,
        organizationID: organizationID,
        search: search,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  listIntegrations(search, applicationID, organizationID, limit, offset, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.ListIntegrations({
        limit: limit,
        offset: offset,
        applicationID: applicationID,
        organizationID: organizationID,
        search: search,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  createHTTPIntegration(integration, callbackFunc, errorCallbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.CreateHTTPIntegration({
        "integration.applicationId": integration.applicationID,
        body: {
          integration: integration,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.integrationNotification("http", "created");
        callbackFunc(resp.obj);
      })
      .catch(error => {
        errorHandler(error);
        if (errorCallbackFunc) errorCallbackFunc(error);
      });
    });
  }

  getHTTPIntegration(applicationID, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.GetHTTPIntegration({
        applicationID: applicationID,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  updateHTTPIntegration(integration, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.UpdateHTTPIntegration({
        "integration.applicationId": integration.applicationID,
        body: {
          integration: integration,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.integrationNotification("http", "updated");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  deleteHTTPIntegration(applicationID, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.DeleteHTTPIntegration({
        applicationID: applicationID,
      })
      .then(checkStatus)
      .then(resp => {
        this.integrationNotification("http", "deleted");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
      ;
    });
  }

  createInfluxDBIntegration(integration, callbackFunc, errorCallbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.CreateInfluxDBIntegration({
        "integration.applicationId": integration.applicationID,
        body: {
          integration: integration,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.integrationNotification("InfluxDB", "created");
        callbackFunc(resp.obj);
      })
      .catch(error => {
        errorHandler(error);
        if (errorCallbackFunc) errorCallbackFunc(error);
      });
    });
  }

  getInfluxDBIntegration(applicationID, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.GetInfluxDBIntegration({
        applicationID: applicationID,
      })
      .then(checkStatus)
      .then(resp => {
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  updateInfluxDBIntegration(integration, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.UpdateInfluxDBIntegration({
        "integration.applicationId": integration.applicationID,
        body: {
          integration: integration,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.integrationNotification("InfluxDB", "updated");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  deleteInfluxDBIntegration(applicationID, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.DeleteInfluxDBIntegration({
        applicationID: applicationID,
      })
      .then(checkStatus)
      .then(resp => {
        this.integrationNotification("InfluxDB", "deleted");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  createThingsBoardIntegration(integration, callbackFunc, errorCallbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.CreateThingsBoardIntegration({
        "integration.applicationId": integration.applicationID,
        body: {
          integration: integration,
        },
      })
        .then(checkStatus)
        .then(resp =>  {
          this.integrationNotification("ThingsBoard.io", "created");
          callbackFunc(resp.obj);
        })
        .catch(error => {
          errorHandler(error);
          if (errorCallbackFunc) errorCallbackFunc(error);
        });
    });
  }

  getThingsBoardIntegration(applicationID, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.GetThingsBoardIntegration({
        applicationID: applicationID,
      })
        .then(checkStatus)
        .then(resp => {
          callbackFunc(resp.obj);
        })
      .catch(errorHandler);
    });
  }

  updateThingsBoardIntegration(integration, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.UpdateThingsBoardIntegration({
        "integration.applicationId": integration.applicationID,
        body: {
          integration: integration,
        },
      })
      .then(checkStatus)
      .then(resp => {
        this.integrationNotification("ThingsBoard.io", "updated");
        callbackFunc(resp.obj);
      })
      .catch(errorHandler);
    });
  }

  deleteThingsBoardIntegration(applicationID, callbackFunc) {
    this.swagger.then(client => {
      client.apis.ApplicationService.DeleteThingsBoardIntegration({
        applicationID: applicationID,
      })
      .then(checkStatus)
      .then(resp => {
        this.integrationNotification("ThingsBoard.io", "deleted");
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
        message: "application has been " + action,
      },
    });
  }

  integrationNotification(kind, action) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "success",
        message: kind + " integration has been " + action,
      },
    });
  }
}

const applicationStore = new ApplicationStore();
export default applicationStore;
