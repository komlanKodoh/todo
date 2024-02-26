import { Url } from './utils';

export type Todo = {
	Id: number;
	Title: string;
	Description: string;
	CompletedAt?: Date;
	CreatedAt: string;
	UpdatedAt: string;
	DeletedAt?: Date;
};

export type CreateTodoDto = {
	Title: string;
	Description: string;
};

export type RequestTodoDto = {
	Title?: string;
	Description?: string;
	Completed?: boolean;
};

export type PatchTodoDto = {
	Id: number;
	Title?: string;
	Description?: string;
	Completed?: boolean;
};

export async function createTodo(todo: CreateTodoDto): Promise<Todo> {
	const response = await fetch(Url.api('/api/todo'), {
		method: 'POST',
		body: JSON.stringify(todo)
	});

	const parsedResponse = await response.json();

	return parsedResponse.Content as Todo;
}

export async function deleteTodo(todoId: number): Promise<Todo> {
	const response = await fetch(Url.api(`/api/todo/${todoId}`), {
		method: 'DELETE'
	});

	const parsedResponse = await response.json();

	return parsedResponse.Content as Todo;
}

export async function getTodos(action: 'get-todos', filters: RequestTodoDto): Promise<Todo[]> {
	const response = await fetch(Url.api('/api/todo/filter'), {
		method: 'POST',
		body: JSON.stringify(filters)
	});

	const parsedResponse = await response.json();

	return parsedResponse.Content as Todo[];
}

export async function updateTodo(todo: PatchTodoDto): Promise<Todo> {
	const response = await fetch(Url.api(`/api/todo/${todo.Id}`), {
		method: 'PATCH',
		body: JSON.stringify(todo)
	});

	const parsedResponse = await response.json();

	return parsedResponse.Content as Todo;
}
