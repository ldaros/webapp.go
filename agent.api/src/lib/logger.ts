export enum LogLevel {
  DEBUG = "DEBUG",
  INFO = "INFO",
  WARN = "WARN",
  ERROR = "ERROR",
}

export class Logger {
  private static instance: Logger;
  private logLevel: LogLevel;

  private constructor() {
    const logLevel = process.env.LOG_LEVEL as LogLevel;
    this.logLevel = this.isValidLogLevel(logLevel) ? logLevel : LogLevel.INFO;
  }

  public static getInstance(): Logger {
    if (!Logger.instance) {
      Logger.instance = new Logger();
    }
    return Logger.instance;
  }

  private isValidLogLevel(logLevel: any): logLevel is LogLevel {
    return Object.values(LogLevel).includes(logLevel);
  }

  private shouldLog(level: LogLevel): boolean {
    const levels: { [key in LogLevel]: number } = {
      [LogLevel.DEBUG]: 0,
      [LogLevel.INFO]: 1,
      [LogLevel.WARN]: 2,
      [LogLevel.ERROR]: 3,
    };

    return levels[level] >= levels[this.logLevel];
  }

  private formatMessage(level: LogLevel, message: string): string {
    const timestamp = new Date().toISOString();
    return `[${timestamp}] [${level.toUpperCase()}] ${message}`;
  }

  public debug(message: string): void {
    if (this.shouldLog(LogLevel.DEBUG)) {
      console.debug(this.formatMessage(LogLevel.DEBUG, message));
    }
  }

  public info(message: string): void {
    if (this.shouldLog(LogLevel.INFO)) {
      console.info(this.formatMessage(LogLevel.INFO, message));
    }
  }

  public warn(message: string): void {
    if (this.shouldLog(LogLevel.WARN)) {
      console.warn(this.formatMessage(LogLevel.WARN, message));
    }
  }

  public error(message: string): void {
    if (this.shouldLog(LogLevel.ERROR)) {
      console.error(this.formatMessage(LogLevel.ERROR, message));
    }
  }
}
