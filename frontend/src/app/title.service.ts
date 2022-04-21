import { Injectable } from '@angular/core';
import { ServiceError, TitleServiceClient } from 'src/generated/service_pb_service';
import { LogEntry, Request } from 'src/generated/service_pb';

@Injectable()
export class TitleService {

  private client: TitleServiceClient;

  constructor() {
    this.client = new TitleServiceClient('http://localhost:8080')
  }

  getTitle(title: string): Promise<string> {
    const request = new Request();
    request.setTitle(title);
    return new Promise<string>((resolve, reject) => {
      this.client.log(request, (error: ServiceError | null, response: LogEntry | null) => {
        if (error !== null || response === null) {
          reject(error);
        }
        resolve(response!.getMessage())
      });
    })
  }
}
