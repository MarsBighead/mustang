from datetime import datetime
#from django.http import HttpResponse
from django.template import loader, Context, RequestContext
from django.shortcuts import render, render_to_response
from django.http import HttpResponseRedirect
from blog.models import BlogPost, BlogPostForm 
# Create your views here.
def archive(request):
    #posts = BlogPost.objects.all()[:10]
    posts = BlogPost.objects.all()
    return render_to_response('archive.html',{'posts':posts, 'form':BlogPostForm}, RequestContext(request))

def create_blogpost(request):
    if request.method == 'POST':
        form = BlogPostForm(request.POST)
        if form.is_valid():
            post = form.save(commit=False)
            post.timestamp = datetime.now()
            post.save()
    return  HttpResponseRedirect('/blog/')

def index(request):
    return HttpResponse("Hello world. You're at the blog index.")
