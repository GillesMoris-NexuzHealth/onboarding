import { Component } from '@angular/core';

import { Title } from './title';
import { TitleService } from './title.service';

@Component({
  selector: 'app-title-form',
  templateUrl: './title-form.component.html',
  styleUrls: ['./title-form.component.css'],
  providers: [ TitleService ]
})
export class TitleFormComponent {

  constructor(private titleService: TitleService) {}

  model = new Title('Frontend');

  async onSubmit() {
    this.model.resultTitle = await this.titleService.getTitle(this.model.title);
  }

}
