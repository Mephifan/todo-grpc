package com.adcubum.syrius.grpc;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

import java.util.ArrayList;

import java.util.Arrays;
import java.util.logging.Logger;

public class TaskServer {

    private static final Logger logger = Logger.getLogger(TaskServer.class.getName());

    private int port = 9090;
    private Server server;

    private Connection connection = null;

    private void start() throws Exception {
        logger.info("Starting the grpc server");

        server = ServerBuilder.forPort(port)
                .addService(new TodoServiceImpl())
                .build()
                .start();

        logger.info("Server started. Listening on port " + port);

        try
        {

            System.out.println("create conection");
            // create a database connection
            connection = DriverManager.getConnection("jdbc:sqlite:data/todos.db");

        }
        catch(SQLException e)
        {
            // if the error message is "out of memory",
            // it probably means no database file is found
            System.err.println(e.getMessage());

        }
        finally
        {}

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.err.println("*** JVM is shutting down. Turning off grpc server as well ***");
            try
            {
                if(connection != null)
                    connection.close();
                System.err.println("*** db disconnected ***");

            }
            catch(SQLException e)
            {
                // connection close failed.
                System.err.println(e.getMessage());
            }
            TaskServer.this.stop();
            System.err.println("*** shutdown complete ***");
        }));
    }

    private void stop() {
        if (server != null) {
            server.shutdown();
        }
    }

    public static void main(String[] args) throws Exception {
        logger.info("Server startup. Args = " + Arrays.toString(args));
        final TaskServer TaskServer = new TaskServer();

        TaskServer.start();
        TaskServer.blockUntilShutdown();
    }

    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    private class TodoServiceImpl extends TodoServiceGrpc.TodoServiceImplBase {

        @Override
        public void listTodo(ListTodoRequest request,
                             io.grpc.stub.StreamObserver<TodoCollection> responseObserver) {

            TodoCollection todos = TodoCollection.newBuilder().build();
            TodoEntity todoEntity;
            Todo todo = null;

            try{

                Statement statement = connection.createStatement();
                statement.setQueryTimeout(30);  // set timeout to 30 sec.

                ResultSet rs = statement.executeQuery("select * from todo limit 23");
                int index = 0;
                while(rs.next())
                {
                    // read the result set
                    todo = generateTodo(rs);

                    //todoEntity = TodoEntity.newBuilder().setData(todo).build();
                    todoEntity = generateTodoEntity(todo);

                    todos = todos.toBuilder().addData( index,todoEntity ).build();
                    index++;

                }
            }
            catch(SQLException e)
            {
                // sql failed.
                System.err.println(e.getMessage());
            }

            responseObserver.onNext(todos);
            responseObserver.onCompleted();
        }

        @Override
        public void getTodo(GetTodoRequest request,
                            io.grpc.stub.StreamObserver<TodoEntity> responseObserver) {

            Todo todo = null;
            TodoEntity response = null;

            try{

                Statement statement = connection.createStatement();
                statement.setQueryTimeout(30);  // set timeout to 30 sec.
                System.out.println("select * from todo where id='"+request.getId()+"'");

                ResultSet rs = statement.executeQuery("select * from todo where id='"+request.getId()+"'");
                while(rs.next())
                {

                    Complete complete = Complete.UNKNOWN;
                    // generate todo entity
                    todo = generateTodo(rs);

                    //response = TodoEntity.newBuilder().setData(todo).build();

                    response = generateTodoEntity(todo);

                    break;
                }
            }
            catch(SQLException e)
            {
                // sql failed.
                System.err.println(e.getMessage());
            }

            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }

    }

    private Todo generateTodo(ResultSet rs) {

        Todo todo = null;

        try{
            todo = Todo.newBuilder().setId(rs.getString("id"))
                    .setTitle(rs.getString("title"))
                    .setDescription(rs.getString("description"))
                    .setCompletedValue(rs.getInt("completed"))
                    .build();
        }
        catch(SQLException e)
        {
            // sql failed.
            System.err.println(e.getMessage());
        }

        return todo;
    }

    private TodoEntity generateTodoEntity(Todo todo) {

        TodoEntity todoEntity = TodoEntity.newBuilder().setData(todo).build();

        Link link = Link.newBuilder().setHref("href").setType("type").build();

        todoEntity = todoEntity.toBuilder()
                .addLinks( 0, Link.newBuilder().setRel("self").setMethodValue(1).setHref("http://localhost:8080/todos/"+todo.getId()).setType("application/json").build())
                .addLinks( 1, Link.newBuilder().setRel("delete").setMethodValue(5).setHref("http://localhost:8080/todos/"+todo.getId()).setType("application/json").build())
                .addLinks( 2, Link.newBuilder().setRel("update").setMethodValue(4).setHref("http://localhost:8080/todos/"+todo.getId()).setType("application/json").build())
                .addLinks( 3, Link.newBuilder().setRel("up").setMethodValue(1).setHref("http://localhost:8080/todos/").setType("application/json").build())
                .build();

        return todoEntity;
    }
}
