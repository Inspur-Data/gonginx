package parser

import (
	"strings"
)

// got the list from: https://nginx.org/en/docs/dirindex.html
var validDirectivesRawList = `absolute_redirect
accept_mutex
accept_mutex_delay
access_by_lua_block
access_log
add_after_body
add_before_body
add_header
add_trailer
addition_types
aio
aio_write
ajp_temp_path
ajp_pass
alias
allow
ancient_browser
ancient_browser_value
api
auth_basic
auth_basic_user_file
auth_delay
auth_digest
auth_digest_user_file
auth_http
auth_http_header
auth_http_pass_client_cert
auth_http_timeout
auth_jwt
auth_jwt_claim_set
auth_jwt_header_set
auth_jwt_key_cache
auth_jwt_key_file
auth_jwt_key_request
auth_jwt_leeway
auth_jwt_require
auth_jwt_type
auth_request
auth_request_set
autoindex
autoindex_exact_size
autoindex_format
autoindex_localtime
balancer_by_lua_block
body_filter_by_lua_block
break
brotli
brotli_comp_level
brotli_min_length
brotli_types
charset
charset_map
charset_types
chunked_transfer_encoding
client_body_buffer_size
client_body_in_file_only
client_body_in_single_buffer
client_body_temp_path
client_body_timeout
client_header_buffer_size
client_header_timeout
client_max_body_size
connect_timeout
connection_pool_size
content_by_lua_block
create_full_put_path
daemon
dav_access
dav_methods
debug_connection
debug_points
default
default_type
deny
directio
directio_alignment
disable_symlinks
empty_gif
env
error_log
error_page
etag
events
expires
f4f
f4f_buffer_size
fastcgi_bind
fastcgi_buffer_size
fastcgi_buffering
fastcgi_buffers
fastcgi_busy_buffers_size
fastcgi_cache
fastcgi_cache_background_update
fastcgi_cache_bypass
fastcgi_cache_key
fastcgi_cache_lock
fastcgi_cache_lock_age
fastcgi_cache_lock_timeout
fastcgi_cache_max_range_offset
fastcgi_cache_methods
fastcgi_cache_min_uses
fastcgi_cache_path
fastcgi_cache_purge
fastcgi_cache_revalidate
fastcgi_cache_use_stale
fastcgi_cache_valid
fastcgi_catch_stderr
fastcgi_connect_timeout
fastcgi_force_ranges
fastcgi_hide_header
fastcgi_ignore_client_abort
fastcgi_ignore_headers
fastcgi_index
fastcgi_intercept_errors
fastcgi_keep_conn
fastcgi_limit_rate
fastcgi_max_temp_file_size
fastcgi_next_upstream
fastcgi_next_upstream_timeout
fastcgi_next_upstream_tries
fastcgi_no_cache
fastcgi_param
fastcgi_pass
fastcgi_pass_header
fastcgi_pass_request_body
fastcgi_pass_request_headers
fastcgi_read_timeout
fastcgi_request_buffering
fastcgi_send_lowat
fastcgi_send_timeout
fastcgi_socket_keepalive
fastcgi_split_path_info
fastcgi_store
fastcgi_store_access
fastcgi_temp_file_write_size
fastcgi_temp_path
flv
geo
geoip2
geoip_city
geoip_country
geoip_org
geoip_proxy
geoip_proxy_recursive
google_perftools_profiles
grpc_bind
grpc_buffer_size
grpc_connect_timeout
grpc_hide_header
grpc_ignore_headers
grpc_intercept_errors
grpc_next_upstream
grpc_next_upstream_timeout
grpc_next_upstream_tries
grpc_pass
grpc_pass_header
grpc_read_timeout
grpc_send_timeout
grpc_set_header
grpc_socket_keepalive
grpc_ssl_certificate
grpc_ssl_certificate_key
grpc_ssl_ciphers
grpc_ssl_conf_command
grpc_ssl_crl
grpc_ssl_name
grpc_ssl_password_file
grpc_ssl_protocols
grpc_ssl_server_name
grpc_ssl_session_reuse
grpc_ssl_trusted_certificate
grpc_ssl_verify
grpc_ssl_verify_depth
gunzip
gunzip_buffers
gzip
gzip_buffers
gzip_comp_level
gzip_disable
gzip_http_version
gzip_min_length
gzip_proxied
gzip_static
gzip_types
gzip_vary
hash
health_check
health_check_timeout
header_filter_by_lua_block
hls
hls_buffers
hls_forward_args
hls_fragment
hls_mp4_buffer_size
hls_mp4_max_buffer_size
http
http2
http2_body_preread_size
http2_chunk_size
http2_idle_timeout
http2_max_concurrent_pushes
http2_max_concurrent_streams
http2_max_field_size
http2_max_header_size
http2_max_requests
http2_push
http2_push_preload
http2_recv_buffer_size
http2_recv_timeout
http3
http3_hq
http3_max_concurrent_streams
http3_stream_buffer_size
if
if_modified_since
ignore_invalid_headers
image_filter
image_filter_buffer
image_filter_interlace
image_filter_jpeg_quality
image_filter_sharpen
image_filter_transparency
image_filter_webp_quality
imap_auth
imap_capabilities
imap_client_buffer
include
index
init_by_lua_block
init_worker_by_lua_block
internal
internal_redirect
ip_hash
js_access
js_body_filter
js_content
js_fetch_buffer_size
js_fetch_ciphers
js_fetch_max_response_buffer_size
js_fetch_protocols
js_fetch_timeout
js_fetch_trusted_certificate
js_fetch_verify
js_fetch_verify_depth
js_filter
js_header_filter
js_import
js_include
js_path
js_periodic
js_preload_object
js_preread
js_set
js_shared_dict_zone
js_var
keepalive
keepalive_disable
keepalive_requests
keepalive_time
keepalive_timeout
keyval
keyval_zone
large_client_header_buffers
least_conn
least_time
limit_conn
limit_conn_dry_run
limit_conn_log_level
limit_conn_status
limit_conn_zone
limit_except
limit_rate
limit_rate_after
limit_req
limit_req_dry_run
limit_req_log_level
limit_req_status
limit_req_zone
limit_zone
lingering_close
lingering_time
lingering_timeout
listen
load_module
location
lock_file
log_by_lua_block
log_format
log_not_found
log_subrequest
lua_add_variable
lua_package_path
lua_shared_dict
mail
map
map_hash_bucket_size
map_hash_max_size
master_process
match
max_errors
max_ranges
memcached_bind
memcached_buffer_size
memcached_connect_timeout
memcached_gzip_flag
memcached_next_upstream
memcached_next_upstream_timeout
memcached_next_upstream_tries
memcached_pass
memcached_read_timeout
memcached_send_timeout
memcached_socket_keepalive
merge_slashes
mgmt
min_delete_depth
mirror
mirror_request_body
modern_browser
modern_browser_value
modsecurity
modsecurity_rules
modsecurity_rules_file
more_clear_headers
more_set_headers
mp4
mp4_buffer_size
mp4_limit_rate
mp4_limit_rate_after
mp4_max_buffer_size
mp4_start_key_frame
mqtt
mqtt_buffers
mqtt_preread
mqtt_rewrite_buffer_size
mqtt_set_connect
msie_padding
msie_refresh
multi_accept
ntlm
open_file_cache
open_file_cache_errors
open_file_cache_min_uses
open_file_cache_valid
open_log_file_cache
opentracing
opentracing_propagate_context
otel_exporter
otel_service_name
otel_span_attr
otel_span_name
otel_trace
otel_trace_context
output_buffers
override_charset
pcre_jit
perl
perl_modules
perl_require
perl_set
pid
pop3_auth
pop3_capabilities
port_in_redirect
postpone_output
preread_buffer_size
preread_by_lua_block
preread_timeout
protocol
proxy_bind
proxy_buffer
proxy_buffer_size
proxy_buffering
proxy_buffers
proxy_busy_buffers_size
proxy_cache
proxy_cache_background_update
proxy_cache_bypass
proxy_cache_convert_head
proxy_cache_key
proxy_cache_lock
proxy_cache_lock_age
proxy_cache_lock_timeout
proxy_cache_max_range_offset
proxy_cache_methods
proxy_cache_min_uses
proxy_cache_path
proxy_cache_purge
proxy_cache_revalidate
proxy_cache_use_stale
proxy_cache_valid
proxy_connect_timeout
proxy_cookie_domain
proxy_cookie_flags
proxy_cookie_path
proxy_download_rate
proxy_force_ranges
proxy_half_close
proxy_headers_hash_bucket_size
proxy_headers_hash_max_size
proxy_hide_header
proxy_http_version
proxy_ignore_client_abort
proxy_ignore_headers
proxy_intercept_errors
proxy_limit_rate
proxy_max_temp_file_size
proxy_method
proxy_next_upstream
proxy_next_upstream_timeout
proxy_next_upstream_tries
proxy_no_cache
proxy_pass
proxy_pass_error_message
proxy_pass_header
proxy_pass_request_body
proxy_pass_request_headers
proxy_protocol
proxy_protocol_timeout
proxy_read_timeout
proxy_redirect
proxy_request_buffering
proxy_requests
proxy_responses
proxy_send_lowat
proxy_send_timeout
proxy_session_drop
proxy_set_body
proxy_set_header
proxy_smtp_auth
proxy_socket_keepalive
proxy_ssl
proxy_ssl_certificate
proxy_ssl_certificate_key
proxy_ssl_ciphers
proxy_ssl_conf_command
proxy_ssl_crl
proxy_ssl_name
proxy_ssl_password_file
proxy_ssl_protocols
proxy_ssl_server_name
proxy_ssl_session_reuse
proxy_ssl_trusted_certificate
proxy_ssl_verify
proxy_ssl_verify_depth
proxy_store
proxy_store_access
proxy_temp_file_write_size
proxy_temp_path
proxy_timeout
proxy_upload_rate
queue
quic_active_connection_id_limit
quic_bpf
quic_gso
quic_host_key
quic_retry
random
random_index
read_ahead
read_timeout
real_ip_header
real_ip_recursive
recursive_error_pages
referer_hash_bucket_size
referer_hash_max_size
request_pool_size
reset_timedout_connection
resolver
resolver_timeout
return
rewrite
rewrite_by_lua_block
rewrite_log
root
satisfy
scgi_bind
scgi_buffer_size
scgi_buffering
scgi_buffers
scgi_busy_buffers_size
scgi_cache
scgi_cache_background_update
scgi_cache_bypass
scgi_cache_key
scgi_cache_lock
scgi_cache_lock_age
scgi_cache_lock_timeout
scgi_cache_max_range_offset
scgi_cache_methods
scgi_cache_min_uses
scgi_cache_path
scgi_cache_purge
scgi_cache_revalidate
scgi_cache_use_stale
scgi_cache_valid
scgi_connect_timeout
scgi_force_ranges
scgi_hide_header
scgi_ignore_client_abort
scgi_ignore_headers
scgi_intercept_errors
scgi_limit_rate
scgi_max_temp_file_size
scgi_next_upstream
scgi_next_upstream_timeout
scgi_next_upstream_tries
scgi_no_cache
scgi_param
scgi_pass
scgi_pass_header
scgi_pass_request_body
scgi_pass_request_headers
scgi_read_timeout
scgi_request_buffering
scgi_send_timeout
scgi_socket_keepalive
scgi_store
scgi_store_access
scgi_temp_file_write_size
scgi_temp_path
secure_link
secure_link_md5
secure_link_secret
send_lowat
send_timeout
sendfile
sendfile_max_chunk
server
server_name
server_name_in_redirect
server_names_hash_bucket_size
server_names_hash_max_size
server_tokens
session_log
session_log_format
session_log_zone
set
set_by_lua_block
set_escape_uri
set_real_ip_from
slice
smtp_auth
smtp_capabilities
smtp_client_buffer
smtp_greeting_delay
source_charset
split_clients
ssi
ssi_last_modified
ssi_min_file_chunk
ssi_silent_errors
ssi_types
ssi_value_length
ssl
ssl_alpn
ssl_buffer_size
ssl_certificate
ssl_certificate_by_lua_block
ssl_certificate_key
ssl_ciphers
ssl_client_certificate
ssl_conf_command
ssl_crl
ssl_dhparam
ssl_early_data
ssl_ecdh_curve
ssl_engine
ssl_handshake_timeout
ssl_name
ssl_ocsp
ssl_ocsp_cache
ssl_ocsp_responder
ssl_password_file
ssl_prefer_server_ciphers
ssl_preread
ssl_protocols
ssl_reject_handshake
ssl_server_name
ssl_session_cache
ssl_session_ticket_key
ssl_session_tickets
ssl_session_timeout
ssl_stapling
ssl_stapling_file
ssl_stapling_responder
ssl_stapling_verify
ssl_trusted_certificate
ssl_verify
ssl_verify_client
ssl_verify_depth
starttls
state
status
status_format
status_zone
sticky
sticky_cookie_insert
stream
stub_status
sub_filter
sub_filter_last_modified
sub_filter_once
sub_filter_types
subrequest_output_buffer_size
tcp_nodelay
tcp_nopush
thread_pool
timeout
timer_resolution
try_files
types
types_hash_bucket_size
types_hash_max_size
underscores_in_headers
uninitialized_variable_warn
upstream
upstream_conf
usage_report
use
user
userid
userid_domain
userid_expires
userid_flags
userid_mark
userid_name
userid_p3p
userid_path
userid_service
uuid_file
uwsgi_bind
uwsgi_buffer_size
uwsgi_buffering
uwsgi_buffers
uwsgi_busy_buffers_size
uwsgi_cache
uwsgi_cache_background_update
uwsgi_cache_bypass
uwsgi_cache_key
uwsgi_cache_lock
uwsgi_cache_lock_age
uwsgi_cache_lock_timeout
uwsgi_cache_max_range_offset
uwsgi_cache_methods
uwsgi_cache_min_uses
uwsgi_cache_path
uwsgi_cache_purge
uwsgi_cache_revalidate
uwsgi_cache_use_stale
uwsgi_cache_valid
uwsgi_connect_timeout
uwsgi_force_ranges
uwsgi_hide_header
uwsgi_ignore_client_abort
uwsgi_ignore_headers
uwsgi_intercept_errors
uwsgi_limit_rate
uwsgi_max_temp_file_size
uwsgi_modifier1
uwsgi_modifier2
uwsgi_next_upstream
uwsgi_next_upstream_timeout
uwsgi_next_upstream_tries
uwsgi_no_cache
uwsgi_param
uwsgi_pass
uwsgi_pass_header
uwsgi_pass_request_body
uwsgi_pass_request_headers
uwsgi_read_timeout
uwsgi_request_buffering
uwsgi_send_timeout
uwsgi_socket_keepalive
uwsgi_ssl_certificate
uwsgi_ssl_certificate_key
uwsgi_ssl_ciphers
uwsgi_ssl_conf_command
uwsgi_ssl_crl
uwsgi_ssl_name
uwsgi_ssl_password_file
uwsgi_ssl_protocols
uwsgi_ssl_server_name
uwsgi_ssl_session_reuse
uwsgi_ssl_trusted_certificate
uwsgi_ssl_verify
uwsgi_ssl_verify_depth
uwsgi_store
uwsgi_store_access
uwsgi_temp_file_write_size
uwsgi_temp_path
valid_referers
variables_hash_bucket_size
variables_hash_max_size
worker_aio_requests
worker_connections
worker_cpu_affinity
worker_priority
worker_processes
worker_rlimit_core
worker_rlimit_nofile
worker_shutdown_timeout
working_directory
xclient
xml_entities
xslt_last_modified
xslt_param
xslt_string_param
xslt_stylesheet
xslt_types
zone
zone_sync
zone_sync_buffers
zone_sync_connect_retry_interval
zone_sync_connect_timeout
zone_sync_interval
zone_sync_recv_buffer_size
zone_sync_server
zone_sync_ssl
zone_sync_ssl_certificate
zone_sync_ssl_certificate_key
zone_sync_ssl_ciphers
zone_sync_ssl_conf_command
zone_sync_ssl_crl
zone_sync_ssl_name
zone_sync_ssl_password_file
zone_sync_ssl_protocols
zone_sync_ssl_server_name
zone_sync_ssl_trusted_certificate
zone_sync_ssl_verify
zone_sync_ssl_verify_depth
zone_sync_timeout
`

// https://github.com/openresty/lua-nginx-module?tab=readme-ov-file#directives
var validLuaDirectivesRawList = `lua_load_resty_core
lua_capture_error_log
lua_use_default_type
lua_malloc_trim
lua_code_cache
lua_thread_cache_max_entries
lua_regex_cache_max_entries
lua_regex_match_limit
lua_package_path
lua_package_cpath
init_by_lua
init_by_lua_block
init_by_lua_file
init_worker_by_lua
init_worker_by_lua_block
init_worker_by_lua_file
exit_worker_by_lua_block
exit_worker_by_lua_file
set_by_lua
set_by_lua_block
set_by_lua_file
content_by_lua
content_by_lua_block
content_by_lua_file
server_rewrite_by_lua_block
server_rewrite_by_lua_file
rewrite_by_lua
rewrite_by_lua_block
rewrite_by_lua_file
access_by_lua
access_by_lua_block
access_by_lua_file
header_filter_by_lua
header_filter_by_lua_block
header_filter_by_lua_file
body_filter_by_lua
body_filter_by_lua_block
body_filter_by_lua_file
log_by_lua
log_by_lua_block
log_by_lua_file
balancer_by_lua_block
balancer_by_lua_file
balancer_keepalive
lua_need_request_body
ssl_client_hello_by_lua_block
ssl_client_hello_by_lua_file
ssl_certificate_by_lua_block
ssl_certificate_by_lua_file
ssl_session_fetch_by_lua_block
ssl_session_fetch_by_lua_file
ssl_session_store_by_lua_block
ssl_session_store_by_lua_file
lua_shared_dict
lua_socket_connect_timeout
lua_socket_send_timeout
lua_socket_send_lowat
lua_socket_read_timeout
lua_socket_buffer_size
lua_socket_pool_size
lua_socket_keepalive_timeout
lua_socket_log_errors
lua_ssl_ciphers
lua_ssl_crl
lua_ssl_protocols
lua_ssl_certificate
lua_ssl_certificate_key
lua_ssl_trusted_certificate
lua_ssl_verify_depth
lua_ssl_conf_command
lua_http10_buffering
rewrite_by_lua_no_postpone
access_by_lua_no_postpone
lua_transform_underscores_in_response_headers
lua_check_client_abort
lua_max_pending_timers
lua_max_running_timers
lua_sa_restart
lua_worker_thread_vm_pool_size
`

// ValidDirectives mapped directives easily find
// todo: this could handle the allowed blocks as well
var ValidDirectives map[string]string = map[string]string{}

func init() {
	validDirective := validDirectivesRawList + validLuaDirectivesRawList
	directives := strings.Split(validDirective, "\n")
	for _, directive := range directives {
		ValidDirectives[strings.TrimSpace(directive)] = strings.TrimSpace(directive)
	}
}
