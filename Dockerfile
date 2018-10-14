FROM plugins/git
ADD git-version /bin/
ENTRYPOINT [ "/bin/git-version" ]
