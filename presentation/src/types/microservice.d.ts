export interface IMicroservice {
    ID: number;
    CreatedAt: Date;
    updatedAt: Date | null;
    DeletedAt: Date | null;
    FriendlyName: string;
    RepoLink: string;
    Status: string;
    UserID: number;
    Inputs: Input[];
    OutputLink: string;
    BackendName: string;
  }