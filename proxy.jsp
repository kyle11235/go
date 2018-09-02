<%@ page pageEncoding="UTF-8"%>
<%@ page contentType="text/html; charset=UTF-8"%>
<%@ page import="java.io.*"%>
<%@ page import="java.net.*"%>
<%@ page import="java.util.*"%>

<%@ page import="org.apache.http.HttpResponse"%>
<%@ page import="org.apache.http.client.HttpClient"%>
<%@ page import="org.apache.http.client.methods.HttpPost"%>
<%@ page import="org.apache.http.entity.StringEntity"%>
<%@ page import="org.apache.http.impl.client.DefaultHttpClient"%>
<%@ page import="org.apache.http.util.EntityUtils"%>

<%
	response.addHeader("Access-Control-Allow-Origin","*");

	String url=request.getParameter("url");
	String data=request.getParameter("data");
	String resultStr = "";


	try {

		HttpClient client = new DefaultHttpClient();
		HttpPost request1 = new HttpPost(url);
		request1.setHeader("Content-Type","application/json");

		StringEntity entity = new StringEntity(d, "UTF-8");
		entity.setContentType("text/xml");
		request1.setEntity(entity);
		HttpResponse response1 = client.execute(request1);

		org.apache.http.HttpEntity httpEntity = response1.getEntity();

		if (httpEntity != null) {
			resultStr = EntityUtils.toString(httpEntity, "UTF-8");
		}

		if (resultStr == null || "".equals(resultStr)) {
			System.err.println("clientResponse is null or empty.");
		}

		System.out.println(resultStr);

	} catch (Exception e) {
		e.printStackTrace();
		resultStr="found error";
	}

	String result = resultStr;
%>
<%=result%>