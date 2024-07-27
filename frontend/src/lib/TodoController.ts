
import {ref} from "vue"
import * as M from "@wailsjs/models"
import {AddTask,GetTasks,RemoveTask,TickTask} from "../../wailsjs/go/main/App"

const TodoController =()=> {
  let _changed = ref(false)
  
  async function list(){
    const res = await GetTasks()
    // return todoList.value;
    console.log(res)
    return res
  }

  async function addTask(content : string){
    // todoList.value.push({content,done:false})
    const result = await AddTask(content)
    invalidate();
    return result
  }
  
  async function removeTask(id : number){
    const result = await RemoveTask(id)
    invalidate() 
    return result
  }

  async function tickTask(id:number,done:boolean){
    const result = await TickTask(done,id)
    invalidate()
    return result
  }

  function invalidate(){
    _changed.value = true
  }

  function validate(){
    _changed.value = false
  }
  
  return {
    addTask,
    removeTask,
    list,
    tickTask,
    invalidate,
    validate,
    asChanged : _changed
  }
}

export interface ITodoController {
  addTask: (content: string) => Promise<M.db.Todo>
  removeTask: (id: number) => Promise<M.db.Todo>;
  list: () => Promise<M.db.Todo[]>;
  tickTask: (id: number,done:boolean) => Promise<M.db.Todo>;
}
export default TodoController;
