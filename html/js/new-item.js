Vue.component('new-item',{
    props:{
        id:Number,
        client:String,
        title:String,
        desc:String,
        price:Number,
        statusArr: Array,
        resultArr: Array,
        statusIndex: Number,
        statusProp: String,
        resultProp:String,
    },
    data:function(){
        return{
            //currentId:null,
            currentClient:null,
            currentTitle:null,
            currentDesc:null,
            currentPrice:null,
            //currentStart:null,
            //currentEnd:null,
            currentStatus:null,
            currentResult:this.resultProp,
        }
    },
    methods:{
        sendInsert:function n(){


            this.currentStatus = this.statusProp;
            this.currentResult = this.resultProp;

            $.ajax({
                url : "/insert",
                type: "POST",

                timeout:5000,

                data:`{
            "client":"`+this.currentClient+`",
            "dealTitle":"`+this.currentTitle+`",
            "dealDesc":"`+this.currentDesc+`",
            "price":`+this.currentPrice+`,
            "status":"`+this.currentStatus+`",
            "result":"`+this.currentResult+`"}`,

                contentType: false,

                cache: false,
                processData:false,


            }).done(function(data, textStatus, jqXHR){ //
                if (jqXHR.status == 201){
//              console.log("result =", jqXHR)
                    $('.modal').modal('hide')
                    app.loading();


                } else {
                    alert("Что-то пошло не так: ")
                }

            }).fail(function(data, textStatus, jqXHR) {
                alert("Не удалось записать: ");
            })
        },
    },
    template: `
      <!-- Modal -->
    <div class="modal fade" :id="'add'+this.statusIndex" tabindex="-1" role="dialog" :aria-labelledby="'label'+this.statusIndex" aria-hidden="true">
          <div class="modal-dialog modal-lg" role="document">
              <div class="modal-content">
                  <div  class="text-center border border-light p-5">
                      <div class="container inputBlock">
                      <div class="row">
                          <div class="col-12">
                                <input type="text" class="form-control mb-4" placeholder="Клиент" name="client" v-model:value="currentClient" >
                          </div>
                      </div>
                      <div class="row">
                          <div class="col-12">
                                <input type="text" class="form-control mb-4" placeholder="Заголовок" name="title" v-model:value="currentTitle">
                          </div>
                      </div>
                      <div class="row">
                          <div class="col-12">
                                <input type="text" class="form-control mb-4" placeholder="Описание" name="desc" v-model:value="currentDesc">
                          </div>
                      </div>
                          <div class="row">
                              <div class="col-12">
                                    <input type="number" class="form-control mb-4" placeholder="0" name="price" v-model:value="currentPrice" maxlength="7">
                              </div>
                          </div>
                      <div class="row">
                          <div class="col-12">
                              <select class="browser-default custom-select mb-4" v-model:value="this.statusProp" name="status">
                                  <option value="" disabled="">Результат</option>
                                  <option v-for="statusList in this.statusArr" :value="statusList" >{{statusList}}</option>
                              </select>
                          </div>
                      </div>
                      <div class="row">
                          <div class="col-12">
                              <select class="browser-default custom-select mb-4 invisiable" v-model="this.resultProp" name="result" >
                                  <option value="" disabled="">Стадия</option>
                                  <option v-for="resultList in this.resultArr" :value="resultList" >{{resultList}}</option>
                              </select>
                          </div>
                      </div>
                      <div class="row">
                            <div class="col-12">
                                  <button type="button" class="btn btn-secondary btn-sm" data-dismiss="modal"><i class="far fa-window-close"></i></button>
                                  <button type="button" class="btn btn-primary btn-sm" v-on:click="this.sendInsert"><i class="far fa-save"></i></button>
                            </div>
                      </div>
                  </div>
              </div>
          </div>
      </div>
              <!-- .Modal -->


    </div>
`});