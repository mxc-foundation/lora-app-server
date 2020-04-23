import { EventEmitter } from "events";
import Swagger from "swagger-client";
import dispatcher from "../dispatcher";
import { checkStatus, errorHandler } from "./helpers";
import sessionStore from "./SessionStore";




class OrganizationStore extends EventEmitter {
  constructor() {
    super();
    this.swagger = new Swagger("/swagger/organization.swagger.json", sessionStore.getClientOpts());
  }

  
  async create(organization) {
    try {
        const client = await this.swagger.then((client) => client);
        let resp = await client.apis.OrganizationService.Create({
          body: {
            organization: organization,
          },
        });
  
        resp = await checkStatus(resp);
        this.emit("create", organization);
        this.notify("created");
        
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  async get(id) {
    try {
        const client = await this.swagger.then((client) => client);
        let resp = await client.apis.OrganizationService.Get({
          id
        });
    
        resp = await checkStatus(resp);
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  async update(organization) {
    try {
        const client = await this.swagger;
        let resp = await client.apis.OrganizationService.Update({
          "organization.id": organization.id,
        body: {
          organization,
        },
        });
  
        resp = await checkStatus(resp);
        this.emit("change", organization);
        this.notify("updated");
        
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  async delete(id) {
    try {
        const client = await this.swagger;
        let resp = await client.apis.OrganizationService.Delete({
          id
        });

        resp = await checkStatus(resp);
        this.emit("delete", id);
        this.notify("deleted");
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  async list(search, limit, offset) {
    try {
        const client = await this.swagger;
        let resp = await client.apis.OrganizationService.List({
          search,
          limit,
          offset,
        });
        
        resp = await checkStatus(resp);
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  async addUser(organizationID, user) {
    try {
        const client = await this.swagger;
        let resp = await client.apis.OrganizationService.AddUser({
          "organizationUser.organizationId": organizationID,
          body: {
            organizationUser: user,
          },
        });
        
        resp = await checkStatus(resp);
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  async getUser(organizationID, userID) {
    try {
        const client = await this.swagger;
        let resp = await client.apis.OrganizationService.GetUser({
          organizationID,
          userID,
        });
        
        resp = await checkStatus(resp);
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  async deleteUser(organizationID, userID) {
    try {
        const client = await this.swagger;
        let resp = await client.apis.OrganizationService.DeleteUser({
          organizationID,
          userID, 
        });

        resp = await checkStatus(resp);
        
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  async updateUser(organizationUser) {
    try {
        const client = await this.swagger;
        let resp = await client.apis.OrganizationService.UpdateUser({
          "organizationUser.organizationId": organizationUser.organizationID,
          "organizationUser.userId": organizationUser.userID,
          body: {
            organizationUser: organizationUser,
          },
        });
  
        resp = await checkStatus(resp);
        
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  async listUsers(organizationId, limit, offset) {
    try {
        const client = await this.swagger;
        let resp = await client.apis.OrganizationService.ListUsers({
          organizationId,
          limit,
          offset,
        });
        
        resp = await checkStatus(resp);
        return resp.obj;
      } catch (error) {
        errorHandler(error);
    }
  }

  notify(action) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "success",
        message: "organization has been " + action,
      },
    });
  }
}

const organizationStore = new OrganizationStore();
export default organizationStore;
