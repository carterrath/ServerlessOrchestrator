import { IUserData } from "./user-data";

export interface IMicroserviceData {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string | null;
    DeletedAt: string | null;
    FriendlyName: string;
    RepoLink: string;
    StatusMessage: string;
    IsActive: boolean;
    User: IUserData;
    Inputs: Input[];
    OutputLink: string;
    BackendName: string;
    ImageID: string;
  }