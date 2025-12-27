// types.ts
export interface DeploymentConfig {
  environment: string;
  region: string;
  clusterName: string;
  serviceName: string;
  containerPort: number;
}

export interface ServiceConfig {
  name: string;
  image: string;
  port: number;
  environment: { [key: string]: string };
}

export interface DeploymentOptions {
  config: DeploymentConfig;
  services: ServiceConfig[];
}

export enum EnvironmentType {
  DEV = 'dev',
  STG = 'stg',
  PROD = 'prod',
}

export enum DeploymentStatus {
  PENDING = 'pending',
  IN_PROGRESS = 'in_progress',
  COMPLETE = 'complete',
  FAILED = 'failed',
}