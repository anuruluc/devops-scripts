// types.ts
export interface DeploymentConfig {
  environment: string;
  region: string;
  instanceType: string;
  amiId: string;
}

export interface ServiceConfig {
  name: string;
  port: number;
  protocol: string;
}

export interface ScriptConfig {
  name: string;
  description: string;
  command: string;
  timeout: number;
}

export type DeploymentStatus = 'pending' | 'in_progress' | 'success' | 'failure';

export interface Deployment {
  id: string;
  config: DeploymentConfig;
  services: ServiceConfig[];
  scripts: ScriptConfig[];
  status: DeploymentStatus;
}

export interface DeploymentLog {
  id: string;
  deploymentId: string;
  timestamp: Date;
  message: string;
  level: 'info' | 'warn' | 'error';
}