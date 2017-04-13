"""toger URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/1.9/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  url(r'^$', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  url(r'^$', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.conf.urls import url, include
    2. Add a URL to urlpatterns:  url(r'^blog/', include('blog.urls'))
"""
from django.conf.urls import include, url
from polls import views

app_name = 'polls'
urlpatterns = [
    # ex: /polls/
    url(r'^$',views.IndexView.as_view(), name= 'index'),
    # ex: /polls/5/
    url(r'^(?P<pk>[0-9]+)/$',views.DetailView.as_view(),  name= 'detail'),
    # ex: /polls/5/results/
    url(r'^(?P<pk>[0-9]+)/results/$',views.ResultsView.as_view(), name= 'results'),
    # ex: /polls/5/vote/
    url(r'^(?P<question_id>[0-9]+)/vote/$',views.vote, name= 'vote'),
]
