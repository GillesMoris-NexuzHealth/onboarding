import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

interface TitleResponse {
    data: string;
}

@Injectable()
export class TitleService {
  constructor(private http: HttpClient) {}

  getTitle(title: string): Observable<TitleResponse> {
    return this.http.post<TitleResponse>('http://localhost:8081', title, {
        observe: 'body',
        responseType: 'json'
    });
  }
}
