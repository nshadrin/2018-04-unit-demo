app = Proc.new do |env|
    colour = "%06x" % (rand * 0xffffff)
    ['200', {
        'Content-Type' => 'text/html',
    }, ["<head><meta http-equiv=\"refresh\" content=1><body><div style='background:##{colour}'>I'm a colorful ruby app: Ruby #{RUBY_VERSION}!</div></body>\n\n"]]
end

run app
