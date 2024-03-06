import { IUserData } from "./user-data";

export interface IMicroserviceData {
    ID: number;
    CreatedAt: Date;
    updatedAt: Date | null;
    DeletedAt: Date | null;
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